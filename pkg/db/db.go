package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func Init(filename string) error {
	var err error
	db, err = sql.Open("sqlite", filename)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scheduler (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date TEXT,
			title TEXT,
			comment TEXT,
			repeat TEXT
		)
	`)
	return err
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
