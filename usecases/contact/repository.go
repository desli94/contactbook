package contact

import (
	ec "github.com/desli94/contactbook/entities/contact"
)

// Repository interface specifies the functions to be implemented by contact repository
type Repository interface {
	Add(c ec.Contact) (ec.Contact, error)
	GetAll() ([]ec.Contact, error)
	GetByID(id int64) (ec.Contact, error)
	Update(id int64, newData ec.Contact) (ec.Contact, error)
	Delete(id int64) (int64, error)
}
