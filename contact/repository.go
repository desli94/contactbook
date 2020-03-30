package contact

import (
	"database/sql"
	"fmt"
)

type repositoryImpl struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) repository {
	return &repositoryImpl{db}
}

func (r *repositoryImpl) add(c ContactDTO) (ContactDTO, error) {
	rows, err := r.db.Query("INSERT INTO contact_cn (firstname_cn, lastname_cn, email_cn, number_cn) VALUES ($1, $2, $3, $4)", c.Firstname, c.Lastname, c.Email, c.Number)

	fmt.Println(*rows)

	if err != nil {
		return ContactDTO{}, err
	}
	return c, nil
}
