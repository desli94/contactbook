package main

import (
	postgres "github.com/desli94/contactbook/infrastructure/postgres"
	restAPI "github.com/desli94/contactbook/infrastructure/rest"

	"github.com/desli94/contactbook/registry"
)

func main() {

	db := postgres.NewDB()

	r := registry.NewRegistry(db)

	restAPI.CreateServer(r)

}
