package main

import (
	"fmt"
	"os"
	"phonebook/contact"
)

func main() {
	db := contact.NewDB()
	cr := contact.NewRepo(db)
	cus := contact.NewUseCase(cr)

	cont, err := cus.Add(contact.ContactDTO{
		Firstname: "Ilsed",
		Lastname:  "M",
		Email:     "ilsed@m.com",
		Number:    "123456789",
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cont)

}
