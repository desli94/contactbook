package main

import (
	"fmt"
	"os"
	ac "phonebook/adapters/contact"

	// ec "phonebook/entities/contact"
	i "phonebook/infrastructure"
	uc "phonebook/usecases/contact"
)

func main() {
	db := i.NewDB()
	contactRepo := ac.NewRepo(db)
	contactInteractor := uc.NewContactInteractor(contactRepo)

	// cont, err := contactInteractor.Add(ec.Contact{
	// 	Firstname: "Ilsed",
	// 	Lastname:  "M",
	// 	Email:     "ilsed@m.com",
	// 	Number:    "123456789",
	// })

	cont, err := contactInteractor.GetAll()

	// cont, err := contactInteractor.GetByID(5)

	// cont, err := contactInteractor.Delete(7)

	// cont, err := contactInteractor.Update(2, ec.Contact{
	// 	Firstname: "Ilsed",
	// 	Lastname:  "M",
	// 	Email:     "desli@m.com",
	// 	Number:    "0987654321",
	// })

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cont)

}
