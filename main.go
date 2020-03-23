package main

func main() {
	contacts := initialize()

	// contacts = contacts.addContact("Desli M")
	// contacts = contacts.addContact("Ilsed M")

	ft := contact{
		firstname: "Garet",
		lastname:  "Bale",
		email:     "garet@bale.com",
		number:    "123451234",
	}

	contacts.addContact(ft)

	// contacts.saveToFile("phonebook.txt")
	contacts.printPhonebook()
}
