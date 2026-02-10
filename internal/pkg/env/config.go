package env

import (
	"os"

	"github.com/joho/godotenv"
)

type ENV string

const (
	PRODUCTION  ENV = "production"
	DEVELOPMENT ENV = "development"
	TEST        ENV = "test"
)

var (
	TODO_PORT   string = "7540"
	MODE        ENV    = "development"
	TODO_DBFILE string = "example.db"
)

func Load() {
	godotenv.Load()

	if os.Getenv("TODO_PORT") != "" {
		TODO_PORT = os.Getenv("TODO_PORT")
	}
	if os.Getenv("ENV") != "" {
		MODE = ENV(os.Getenv("ENV"))
	}
	if os.Getenv("TODO_DBFILE") != "" {
		TODO_DBFILE = os.Getenv("TODO_DBFILE")
	}
}
