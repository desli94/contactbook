package main

import (
	"fmt"
	"os"
	"phonebook/contact"
)

func main() {
	// contacts := initialize()

	// ft := contact{
	// 	firstname: "Garet",
	// 	lastname:  "Bale",
	// 	email:     "garet@bale.com",
	// 	number:    "123451234",
	// }

	// contacts.addContact(ft)

	// // contacts.saveToFile("phonebook.txt")
	// contacts.printPhonebook()

	// contactlist.Open()
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
