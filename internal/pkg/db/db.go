package db

import (
	"database/sql"
	"os"

	"diploma/internal/pkg/logger"

	_ "embed"

	_ "modernc.org/sqlite"
)

var Database *sql.DB

//go:embed data.sql
var dataSQL []byte

func Init(path string) error {
	_, err := os.Stat(path)

	if err != nil {
		file, err := os.Create(path)
		if err != nil {
			logger.Get().Error("Couldn't create database file")
			return err
		}

		if err = file.Close(); err != nil {
			logger.Get().Error("Couldn't close database file")
			return err
		}
	}

	Database, err = sql.Open("sqlite", path)

	if err != nil {
		logger.Get().Error("SQL database file not found")
		return err
	}
	logger.Get().Info("Database connection opened")

	_, err = Database.Exec(string(dataSQL))

	if err != nil {
		logger.Get().Error("Couldn't sync data.sql")
		return err
	}

	logger.Get().Info("Database data synced")

	return nil
}

func Close() error {
	if Database == nil {
		return nil
	}

	logger.Get().Info("Database connection closed")
	return Database.Close()
}
