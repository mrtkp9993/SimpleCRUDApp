package main

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func (a *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			respondWithError(w, http.StatusUnauthorized, "Invalid/Missing Credentials.")
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid/Missing Credentials.")
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			respondWithError(w, http.StatusUnauthorized, "Invalid/Missing Credentials.")
			return
		}

		user := User{Username: pair[0]}
		row := a.DB.QueryRow("SELECT id, saltedpassword, salt FROM users WHERE username=?", user.Username)
		if err := row.Scan(&user.Id, &user.Saltedpassword, &user.Salt); err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid/Missing Credentials.")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Saltedpassword), []byte(pair[1]+user.Salt)); err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid/Missing Credentials.")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *App) cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			next.ServeHTTP(w, r)
			return
		}

		content, err := a.Cache.Get(r.RequestURI).Result()
		if err != nil {
			rr := httptest.NewRecorder()
			next.ServeHTTP(rr, r)
			content = rr.Body.String()
			err = a.Cache.Set(r.RequestURI, content, 10*time.Minute).Err()
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}
			respondWithString(w, http.StatusOK, content)
			return
		} else {
			respondWithString(w, http.StatusOK, content)
			return
		}
	})
}
