package app

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

func Greet() {
	// TOMLファイルを読み込み
	content, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}

	// TOMLファイルを構造体にパース
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		panic(err)
	}

	// パース結果を表示
	fmt.Println("Title:", config.Title)
}
