package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb(url string) error {
	if db != nil {
		return fmt.Errorf("the database is already initialized")
	}

	newDb, err := sql.Open("mysql", "user:password@localhost:3306/db")
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	db = newDb

	return nil
}

func getDb() *sql.DB {
	return db
}
