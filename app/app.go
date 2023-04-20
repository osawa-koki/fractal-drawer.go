package app

import (
	"fmt"
  "path/filepath"
	"github.com/BurntSushi/toml"
)

func App() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println("Failed to decode file: ", err)
		return
	}

	config.Mandelbrot.Width = config.Global.Width
	config.Mandelbrot.Height = config.Global.Height
	config.Mandelbrot.OutputFile = filepath.Join(config.Global.OutputDirectory, config.Mandelbrot.OutputFile)
	MandelbrotDrawer(config.Mandelbrot)

	config.Julia.Width = config.Global.Width
	config.Julia.Height = config.Global.Height
	config.Julia.OutputFile = filepath.Join(config.Global.OutputDirectory, config.Julia.OutputFile)
	JuliaDrawer(config.Julia)
}
