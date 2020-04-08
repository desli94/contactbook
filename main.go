package main

import (
	ac "github.com/desli94/contactbook/adapters/contact"
	postgres "github.com/desli94/contactbook/infrastructure/postgres"
	restAPI "github.com/desli94/contactbook/infrastructure/rest"
	uc "github.com/desli94/contactbook/usecases/contact"
)

func main() {

	db := postgres.NewDB()
	contactRepo := ac.NewRepo(db)
	contactInteractor := uc.NewContactInteractor(contactRepo)
	contactController := ac.NewContactController(contactInteractor)

	restAPI.CreateServer(contactController)

}
