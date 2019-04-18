/*
 * Copyright (C) 2019  Murat Koptur
 *
 * Contact: mkoptur3@gmail.com
 *
 * Last edit: 3/31/19 6:03 PM
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

package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"SimpleCRUDApp"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(viper.GetString("DB.username"),
		viper.GetString("DB.password"),
		viper.GetString("DB.server"),
		viper.GetString("DB.port"),
		viper.GetString("DB.db_name"))

	createTestTable()
	createTestData()

	code := m.Run()

	clearTestTable()

	os.Exit(code)
}

func createTestTable() {
	query := `create table if not exists products_test
	(
		id           int(11) unsigned auto_increment primary key,
		name         tinytext null,
		manufacturer tinytext null
	);`
	if _, err := a.DB.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createTestData() {
	query := `INSERT INTO products_test (id, name, manufacturer) VALUES 
            (1, 'eum', 'Osinski-Hagenes'),
			(2, 'eos', 'Runolfsson-Jacobi'),
			(3, 'incidunt', 'Mosciski Ltd')`
	if _, err := a.DB.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func clearTestTable() {
	query := `DROP TABLE products_test`
	if _, err := a.DB.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func TestGetProduct(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/products/2", nil)
	request.SetBasicAuth("user1", "pass1")
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	assert.Equal(t, response.Code, http.StatusOK)
}

func TestCreateProduct(t *testing.T) {
	newProduct := main.Product{Id: 1, Name: "Earthquake Pills", Manufacturer: "ACME"}
	payload, _ := json.Marshal(newProduct)
	request, _ := http.NewRequest("POST", "/api/products/new", bytes.NewReader(payload))
	request.SetBasicAuth("user1", "pass1")
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	assert.Equal(t, response.Code, http.StatusCreated)
}

func TestUpdateProduct(t *testing.T) {
	updatedProduct := main.Product{Id: 1, Name: "Do-It-Yourself Tornado Kit", Manufacturer: "ACME"}
	payload, _ := json.Marshal(updatedProduct)
	request, _ := http.NewRequest("PUT", "/api/products/1", bytes.NewReader(payload))
	request.SetBasicAuth("user1", "pass1")
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	assert.Equal(t, response.Code, http.StatusOK)

	responseData := main.Product{}
	responseBody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseBody, &responseData)
	assert.Equal(t, responseData.Id, 1)
}

func TestDeleteProduct(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/api/products/1", nil)
	request.SetBasicAuth("user1", "pass1")
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	assert.Equal(t, response.Code, http.StatusOK)
}
