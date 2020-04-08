package postgres

import (
	"database/sql"
	"fmt"

	// Third party library to handle database connection
	_ "github.com/lib/pq"
)

// NewDB opens a db connection
func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/phonebook?sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db")
	return db
}
