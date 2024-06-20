package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	// Connect to a sqlite3 database
	db, err := sql.Open("sqlite3", "fiadores.db")
	if err != nil {
		panic(err)
	}
	return db
}
