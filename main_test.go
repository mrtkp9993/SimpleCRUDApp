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
	"github.com/spf13/viper"
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

	os.Exit(m.Run())
}

// TODO Write tests
func TestGetProducts(t *testing.T) {

}

func TestGetProduct(t *testing.T) {

}

func TestCreateProduct(t *testing.T) {

}

func TestUpdateProduct(t *testing.T) {

}

func TestDeleteProduct(t *testing.T) {

}
