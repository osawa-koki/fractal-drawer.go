package app

type Config struct {
	Global GlobalConfig `toml:"global"`
}

type GlobalConfig struct {
	Width           int    `toml:"width"`
	Height          int    `toml:"height"`
	OutputDirectory string `toml:"output_directory"`
}
