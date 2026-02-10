package db

import (
	"database/sql"
	"diploma/internal/pkg/logger"
	_ "embed"
	"os"

	_ "modernc.org/sqlite"
)

//go:embed data.sql
var dataSQL []byte

func Init(path string) error {
	_, err := os.Stat(path)

	if err != nil {
		_, err = os.Create(path)
	}

	db, err := sql.Open("sqlite", path)

	if err != nil {
		logger.Get().Error("SQL database file not found")
		return err
	}
	defer db.Close()

	_, err = db.Exec(string(dataSQL))

	if err != nil {
		logger.Get().Error("Couldn't sync data.sql")
		return err
	}

	logger.Get().Info("Database synced")

	return nil
}
