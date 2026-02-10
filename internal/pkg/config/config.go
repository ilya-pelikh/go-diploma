package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ENV string

const (
	PRODUCTION  ENV = "production"
	DEVELOPMENT ENV = "development"
	TEST        ENV = "test"
)

type Config struct {
	TODO_PORT   string
	MODE        ENV
	TODO_DBFILE string
}

var environment = Config{
	TODO_PORT:   "7540",
	MODE:        "development",
	TODO_DBFILE: "example.db",
}

func Load() {
	godotenv.Load()

	if os.Getenv("TODO_PORT") != "" {
		environment.TODO_PORT = os.Getenv("TODO_PORT")
	}
	if os.Getenv("ENV") != "" {
		environment.MODE = ENV(os.Getenv("ENV"))
	}
	if os.Getenv("TODO_DBFILE") != "" {
		environment.TODO_DBFILE = os.Getenv("TODO_DBFILE")
	}

	fmt.Printf("Configuration:\n%+v\n", environment)

}

func Get() *Config {
	return &environment
}
