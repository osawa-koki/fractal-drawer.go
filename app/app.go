package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

func App() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println("Failed to decode file: ", err)
		return
	}

	// ファイル名を格納するリストを定義
	var files []string

	config.Mandelbrot.Width = config.Global.Width
	config.Mandelbrot.Height = config.Global.Height
	files = append(files, config.Mandelbrot.OutputFile)
	config.Mandelbrot.OutputFile = filepath.Join(config.Global.OutputDirectory, config.Mandelbrot.OutputFile)
	MandelbrotDrawer(config.Mandelbrot)

	config.Julia.Width = config.Global.Width
	config.Julia.Height = config.Global.Height
	files = append(files, config.Julia.OutputFile)
	config.Julia.OutputFile = filepath.Join(config.Global.OutputDirectory, config.Julia.OutputFile)
	JuliaDrawer(config.Julia)

	config.Tricorn.Width = config.Global.Width
	config.Tricorn.Height = config.Global.Height
	files = append(files, config.Tricorn.OutputFile)
	config.Tricorn.OutputFile = filepath.Join(config.Global.OutputDirectory, config.Tricorn.OutputFile)
	TricornDrawer(config.Tricorn)

	config.BurningShip.Width = config.Global.Width
	config.BurningShip.Height = config.Global.Height
	files = append(files, config.BurningShip.OutputFile)
	config.BurningShip.OutputFile = filepath.Join(config.Global.OutputDirectory, config.BurningShip.OutputFile)
	BurningShipDrawer(config.BurningShip)

	{ /* ファイル名一覧を出力 */
		// 新しいファイルのパス
		outputFile := filepath.Join(config.Global.OutputDirectory, "items.txt")

		// 新しいファイルを作成
		out, err := os.Create(outputFile)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		// スライスの要素を改行で結合して新しいファイルに書き込む
		line := strings.Join(files, "\n")
		_, err = out.WriteString(line)
		if err != nil {
			panic(err)
		}

		if err := out.Sync(); err != nil {
			panic(err)
		}
	}
}
