package contact

import (
	ec "phonebook/entities/contact"
)

type contactInteractor struct {
	Repository ContactRepository
}

type ContactDTO struct {
	Firstname string
	Lastname  string
	Email     string
	Number    string
}

func NewContactInteractor(rep ContactRepository) contactInteractor {
	return contactInteractor{rep}
}

func (ci contactInteractor) Add(c ContactDTO) (ContactDTO, error) {
	cont, err := ec.MakeNewContact(ec.Contact{c.Firstname, c.Lastname, c.Email, c.Number})

	if err != nil {
		return ContactDTO{}, err
	}
	contDTO, err := ci.Repository.Add(ContactDTO{cont.Firstname, cont.Lastname, cont.Email, cont.Number})

	if err != nil {
		return ContactDTO{}, err
	}

	return contDTO, nil
}
