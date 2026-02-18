package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type env struct {
	TODO_PORT     string
	TODO_ENV      string
	TODO_DBFILE   string
	TODO_PASSWORD string
	TODO_JWT_KEY  string
}

var ENV *env

func Load() {
	config, err := godotenv.Read()
	if err != nil {
		fmt.Println("Couldn't read config from env")
		return
	}

	ENV = &env{
		TODO_PORT:     config["TODO_PORT"],
		TODO_ENV:      config["TODO_ENV"],
		TODO_DBFILE:   config["TODO_DBFILE"],
		TODO_PASSWORD: config["TODO_PASSWORD"],
		TODO_JWT_KEY:  config["TODO_JWT_KEY"],
	}
}
