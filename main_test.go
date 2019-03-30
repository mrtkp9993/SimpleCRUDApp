/*
 * Copyright (C) 2019  muratkoptur
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
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
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

// TODO Write tests
func TestGetProduct(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/products/1", nil)
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreateProduct(t *testing.T) {

}

func TestUpdateProduct(t *testing.T) {

}

func TestDeleteProduct(t *testing.T) {

}
