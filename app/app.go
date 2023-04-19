package app

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func Greet() {
	var c Configuration
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		fmt.Println("Failed to decode file : ", err)
		return
	}
	fmt.Println("DB username: ", c.DB.Username)
	fmt.Println("DB password: ", c.DB.Password)
	fmt.Println("DB name: ", c.DB.DBName)
}
