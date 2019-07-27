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

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

var err error

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	a := App{}
	a.Initialize(viper.GetString("DB.username"),
		viper.GetString("DB.password"),
		viper.GetString("DB.server"),
		viper.GetString("DB.port"),
		viper.GetString("DB.db_name"),
		viper.GetString("Cache.addr"),
		viper.GetString("Cache.password"))

	a.Run(viper.GetString("Server.port"))
}
