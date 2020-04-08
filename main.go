package main

import (
	// "fmt"
	// "os"

	// ac "github.com/desli94/contactbook/adapters/contact"

	// // ec "github.com/desli94/contactbook/entities/contact"
	// i "github.com/desli94/contactbook/infrastructure"
	ac "github.com/desli94/contactbook/adapters/contact"
	i "github.com/desli94/contactbook/infrastructure"
	uc "github.com/desli94/contactbook/usecases/contact"
)

func main() {
	// db := i.NewDB()
	// contactRepo := ac.NewRepo(db)
	// contactInteractor := uc.NewContactInteractor(contactRepo)

	// cont, err := contactInteractor.Add(ec.Contact{
	// 	Firstname: "Ilsed",
	// 	Lastname:  "M",
	// 	Email:     "ilsed@m.com",
	// 	Number:    "123456789",
	// })

	// cont, err := contactInteractor.GetAll()

	// cont, err := contactInteractor.GetByID(5)

	// cont, err := contactInteractor.Delete(7)

	// cont, err := contactInteractor.Update(2, ec.Contact{
	// 	Firstname: "Ilsed",
	// 	Lastname:  "M",
	// 	Email:     "desli@m.com",
	// 	Number:    "0987654321",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(cont)

	db := i.NewDB()
	contactRepo := ac.NewRepo(db)
	contactInteractor := uc.NewContactInteractor(contactRepo)
	contactController := ac.NewContactController(contactInteractor)

	i.CreateServer(contactController)

}
