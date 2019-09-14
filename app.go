/*
 * Copyright (C) 2019  Murat Koptur
 *
 * Contact: mkoptur3@gmail.com
 *
 * Last edit: 4/1/19 12:23 PM
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	Logger http.Handler
	DB     *sql.DB
	Cache  *redis.Client
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/api/products/list", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/api/products/new", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/api/products/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/api/products/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/api/products/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *App) Initialize(username, password, server, port, dbName, cacheAddr, cachePass string) {
	dataSource := username + ":" + password + "@tcp(" + server + ":" + port + ")/" + dbName
	a.DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	a.Cache = redis.NewClient(&redis.Options{
		Addr:     cacheAddr,
		Password: cachePass,
		DB:       0,
	})

	a.Router = mux.NewRouter()
	a.Logger = handlers.CombinedLoggingHandler(os.Stdout, a.Router)
	a.Router.Use(a.authMiddleware, a.cacheMiddleware)
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	// https://stackoverflow.com/questions/38376226/how-to-allow-options-method-from-mobile-using-gorilla-handler
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+viper.GetString("Server.port"),
		handlers.CORS(headersOk, originsOk, methodsOk)(a.Logger)))
}
