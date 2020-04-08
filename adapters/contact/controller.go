package contact

import (
	ec "github.com/desli94/contactbook/entities/contact"
	uc "github.com/desli94/contactbook/usecases/contact"
)

type controller struct {
	Interactor uc.Interactor
}

// Controller interface specifies the methods to be called by infrastructure
type Controller interface {
	Add(firstname string, lastname string, email string, number string) (ec.Contact, error)
	GetAll() ([]ec.Contact, error)
}

var contactRepo uc.Repository

// NewContactController creates the controller with the injected dependecies
func NewContactController(ci uc.Interactor) Controller {
	return &controller{ci}
}

func (c *controller) Add(firstname string, lastname string, email string, number string) (ec.Contact, error) {
	var contact ec.Contact

	contact = ec.Contact{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Number:    number,
	}

	addedContact, err := c.Interactor.Add(contact)

	if err != nil {
		return addedContact, err
	}

	return addedContact, nil
}

func (c *controller) GetAll() ([]ec.Contact, error) {
	contacts, err := c.Interactor.GetAll()

	if err != nil {
		return contacts, err
	}

	return contacts, nil
}
