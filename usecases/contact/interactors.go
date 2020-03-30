package contact

import (
	"errors"
	ec "phonebook/entities/contact"
)

type contactInteractor struct {
	Repository ContactRepository
}

func NewContactInteractor(rep ContactRepository) contactInteractor {
	return contactInteractor{rep}
}

func (ci contactInteractor) Add(c ContactDTO) (ContactDTO, error) {
	cont, err := ec.MakeNewContact(ec.Contact{c.Firstname, c.Lastname, c.Email, c.Number})

	if err != nil {
		return ContactDTO{}, err
	}
	contDTO, err := ci.Repository.Add(ContactDTO{0, cont.Firstname, cont.Lastname, cont.Email, cont.Number})

	if err != nil {
		return ContactDTO{}, err
	}

	return contDTO, nil
}

func (ci contactInteractor) GetById(id int64) (ContactDTO, error) {
	contact, err := ci.Repository.GetById(id)

	if err != nil {
		return ContactDTO{}, err
	}

	return contact, nil
}

func (ci *contactInteractor) Update(id int64, newData ContactDTO) (ContactDTO, error) {
	updatedContact, err := ci.Repository.Update(id, newData)

	if err != nil {
		return ContactDTO{}, err
	}

	return updatedContact, nil
}

func (ci contactInteractor) Delete(id int64) (int64, error) {
	deletedContacts, err := ci.Repository.Delete(id)

	if err != nil {
		return 0, err
	}
	if deletedContacts == 0 {
		return 0, errors.New("Err: Not deleted")
	}

	return deletedContacts, nil
}
