package contact

import (
	"database/sql"
	"fmt"
	uc "phonebook/usecases/contact"
)

type repositoryImpl struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) uc.ContactRepository {
	return &repositoryImpl{db}
}

func (r *repositoryImpl) Add(c uc.ContactDTO) (uc.ContactDTO, error) {
	rows, err := r.db.Query("INSERT INTO contact_cn (firstname_cn, lastname_cn, email_cn, number_cn) VALUES ($1, $2, $3, $4)", c.Firstname, c.Lastname, c.Email, c.Number)

	fmt.Println(*rows)

	if err != nil {
		return uc.ContactDTO{}, err
	}
	return c, nil
}
