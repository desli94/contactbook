package contact

import (
	"errors"

	ec "github.com/desli94/contactbook/entities/contact"
)

type interactor struct {
	Repository Repository
}

// Interactor interface specifies the methods to be called by controller
type Interactor interface {
	Add(c ec.Contact) (ec.Contact, error)
	GetAll() ([]ec.Contact, error)
	GetByID(id int64) (ec.Contact, error)
	Update(id int64, newData ec.Contact) (ec.Contact, error)
	Delete(id int64) (int64, error)
}

// NewContactInteractor creates the interactor with the injected dependecies
func NewContactInteractor(rep Repository) Interactor {
	return &interactor{rep}
}

func (ci *interactor) Add(c ec.Contact) (ec.Contact, error) {
	contact, err := ec.MakeNewContact(c)

	if err != nil {
		return contact, err
	}
	contact, err = ci.Repository.Add(c)

	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (ci *interactor) GetAll() ([]ec.Contact, error) {
	contacts, err := ci.Repository.GetAll()

	if err != nil {
		return contacts, err
	}

	return contacts, nil
}

func (ci *interactor) GetByID(id int64) (ec.Contact, error) {
	contact, err := ci.Repository.GetByID(id)

	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (ci *interactor) Update(id int64, newData ec.Contact) (ec.Contact, error) {
	updatedContact, err := ci.Repository.Update(id, newData)

	if err != nil {
		return updatedContact, err
	}

	return updatedContact, nil
}

func (ci *interactor) Delete(id int64) (int64, error) {
	deletedContacts, err := ci.Repository.Delete(id)

	if err != nil {
		return 0, err
	}
	if deletedContacts == 0 {
		return 0, errors.New("Err: Not deleted")
	}

	return deletedContacts, nil
}
