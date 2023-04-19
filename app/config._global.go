package app

type ConfigGlobal struct {
	Width           int `toml:"width"`
	Height          int `toml:"height"`
	OutputDirectory string `toml:"output_directory"`
}
