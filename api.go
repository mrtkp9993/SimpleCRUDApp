package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// curl --user user1:pass1 127.0.0.1:8000/api/products/list
func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := a.DB.Query("SELECT * FROM products")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Manufacturer)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}

		products = append(products, p)
	}

	_ = json.NewEncoder(w).Encode(products)
}

// curl --header "Content-Type: application/json" --request POST --data '{"name": "ABC", "manufacturer": "ACME"}' \
// 		--user user1:pass1 127.0.0.1:8000/api/products/new
func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	_, err := a.DB.Query("INSERT INTO products (name, manufacturer) VALUES (?, ?)", p.Name, p.Manufacturer)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithMessage(w, http.StatusCreated, "New row added.")
}

// curl --user user1:pass1 127.0.0.1:8000/api/products/10
func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := Product{Id: id}
	row := a.DB.QueryRow("SELECT name, manufacturer FROM products WHERE id=?", p.Id)
	if err := row.Scan(&p.Name, &p.Manufacturer); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

// curl --request PUT --data '{"name": "ABC", "manufacturer": "ACME"}' --user user1:pass1 127.0.0.1:8000/api/products/11
func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()
	p.Id = id

	_, err = a.DB.Query("UPDATE products SET name=?, manufacturer=? WHERE id=?", p.Name, p.Manufacturer, p.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

// curl --request DELETE --user user1:pass1 127.0.0.1:8000/api/products/10
func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	_, err = a.DB.Query("DELETE FROM products WHERE id=?", id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithMessage(w, http.StatusOK, "Deleted.")
}
