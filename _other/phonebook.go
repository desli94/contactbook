package other

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type contact struct {
	firstname string
	lastname  string
	email     string
	number    string
}

type phonebook []contact

func initialize() phonebook {
	phonebook := phonebook{}
	contact := contact{}
	bs, err := ioutil.ReadFile("phonebook.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	phonebookArr := strings.Split(string(bs), "\n")
	phonebookArr = phonebookArr[:len(phonebookArr)-1]

	for i := range phonebookArr {
		_contact := strings.Split(phonebookArr[i], ",")

		contact.firstname = _contact[0]
		contact.lastname = _contact[1]
		contact.email = _contact[2]
		contact.number = _contact[3]
		phonebook = append(phonebook, contact)
	}
	return phonebook
}

func (p *phonebook) addContact(contact contact) {
	*p = append(*p, contact)
	(*p).saveToFile("phonebook.txt")
}

// Any variable of type phonebook gets access to the printPhonebook method
// The actual copy of phonebook we are working with is available in the function as a varaible callled p
func (p phonebook) printPhonebook() {
	for i, contact := range p {
		fmt.Println("Contact", i+1)
		fmt.Println("Firstname:", contact.firstname)
		fmt.Println("Lastname:", contact.lastname)
		fmt.Println("Email:", contact.email)
		fmt.Println("Number:", contact.number)
		fmt.Println()
	}
}

func (p phonebook) toString() string {
	var contactStr string
	var phonebookStr string
	for _, contact := range p {
		contactStr = contact.firstname + "," + contact.lastname + "," + contact.email + "," + contact.number
		phonebookStr += contactStr + "\n"
	}
	return phonebookStr
}

func (p phonebook) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(p.toString()), 0666)
}
