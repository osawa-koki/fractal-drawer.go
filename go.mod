module example.com/main

go 1.20

replace example.com/main/app => ./app

require example.com/main/app v0.0.0-00010101000000-000000000000

require github.com/BurntSushi/toml v1.2.1 // indirect
