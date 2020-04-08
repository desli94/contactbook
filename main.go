package main

import (
	ac "github.com/desli94/contactbook/adapters/contact"
	i "github.com/desli94/contactbook/infrastructure"
	uc "github.com/desli94/contactbook/usecases/contact"
)

func main() {

	db := i.NewDB()
	contactRepo := ac.NewRepo(db)
	contactInteractor := uc.NewContactInteractor(contactRepo)
	contactController := ac.NewContactController(contactInteractor)

	i.CreateServer(contactController)

}
