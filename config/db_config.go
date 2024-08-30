package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)

	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal("db is null")
	}

	createTable(db)

	return db
}

func createTable(db *sql.DB) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	);`

	_, err := db.Exec(sqlTable)

	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
