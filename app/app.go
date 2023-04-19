package app

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func Greet() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println("Failed to decode file : ", err)
		return
	}
	fmt.Println("Width: ", config.Global.Width)
	fmt.Println("Height: ", config.Global.Height)
	fmt.Println("Output Directory: ", config.Global.OutputDirectory)
}
