package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Database *sql.DB
}

func (db *Database) Execute(statement string, args ...interface{}) {
	var err error
	if len(args) == 0 {
		_, err = db.Database.Exec(statement)
	} else {
		_, err = db.Database.Exec(statement, args...)
	}
	checkError(err, "Executing database statement", true)
}

func (db *Database) Initialize() {
	var err error
	db.Database, err = sql.Open("sqlite3", "build/database.db")
	checkError(err, "Opening database", true)

	statement := "CREATE TABLE IF NOT EXISTS codes(id INTEGER PRIMARY KEY AUTOINCREMENT, hash text NOT NULL, url text NOT NULL, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)"
	db.Execute(statement, NULL)
}

func (db *Database) StoreHash(hashed, url string) {
	db.Execute("INSERT INTO codes(hash, url) VALUES(?, ?)", hashed, url)
}
