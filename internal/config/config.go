package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TODO_PORT string
}

var config Config

func Get() *Config {
	return &config
}

func Load() {
	godotenv.Load()

	config = Config{TODO_PORT: "7540"}

	if os.Getenv("TODO_PORT") != "" {
		config.TODO_PORT = os.Getenv("TODO_PORT")
	}

	fmt.Printf("%+v\n", config)
}
