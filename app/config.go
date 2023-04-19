package app

type Configuration struct {
	DB DBConfiguration
}

type DBConfiguration struct {
	Username string `toml:"user"`
	Password string `toml:"pass"`
	DBName   string `toml:"db"`
}
