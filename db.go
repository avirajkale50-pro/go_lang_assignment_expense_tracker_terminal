package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./expenses.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL,
		category TEXT,
		note TEXT,
		date TEXT
	);
	`
	_, err = db.Exec(query)
	return err
}
