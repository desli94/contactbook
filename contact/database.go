package contact

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/phonebook?sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db")
	return db
}
