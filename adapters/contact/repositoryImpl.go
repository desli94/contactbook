package contact

import (
	"database/sql"
	"fmt"
	ec "github.com/desli94/contactbook/entities/contact"
	uc "github.com/desli94/contactbook/usecases/contact"
)

type repositoryImpl struct {
	db *sql.DB
}

// NewRepo creates a reposotry which hold the database instance.
func NewRepo(db *sql.DB) uc.Repository {
	return &repositoryImpl{db}
}

func (r *repositoryImpl) Add(c ec.Contact) (ec.Contact, error) {
	rows := r.db.QueryRow("INSERT INTO contact_cn (firstname_cn, lastname_cn, email_cn, number_cn) VALUES ($1, $2, $3, $4) RETURNING *", c.Firstname, c.Lastname, c.Email, c.Number)

	contact := ec.Contact{}
	err := rows.Scan(&contact.ID, &contact.Firstname, &contact.Lastname, &contact.Email, &contact.Number)

	if err != nil {
		fmt.Println("yep")
		return contact, err
	}

	return contact, nil
}

func (r *repositoryImpl) GetAll() ([]ec.Contact, error) {
	rows, err := r.db.Query("SELECT * FROM contact_cn ORDER BY id_cn")

	if err != nil {
		return []ec.Contact{}, err
	}

	contacts := []ec.Contact{}

	for rows.Next() {
		contact := ec.Contact{}
		err = rows.Scan(&contact.ID, &contact.Firstname, &contact.Lastname, &contact.Email, &contact.Number)

		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *repositoryImpl) GetByID(id int64) (ec.Contact, error) {
	rows, err := r.db.Query("SELECT * FROM contact_cn WHERE id_cn=$1", id)

	if err != nil {
		return ec.Contact{}, err
	}

	contact := ec.Contact{}
	for rows.Next() {
		err = rows.Scan(&contact.ID, &contact.Firstname, &contact.Lastname, &contact.Email, &contact.Number)

		if err != nil {
			return contact, err
		}
	}

	return contact, nil
}

func (r *repositoryImpl) Update(id int64, newData ec.Contact) (ec.Contact, error) {
	rows := r.db.QueryRow("UPDATE contact_cn SET firstname_cn = $1, lastname_cn = $2, email_cn = $3, number_cn = $4 WHERE id_cn = $5 RETURNING *",
		newData.Firstname, newData.Lastname, newData.Email, newData.Number, id)

	updatedContact := ec.Contact{}
	err := rows.Scan(&updatedContact.ID, &updatedContact.Firstname, &updatedContact.Lastname, &updatedContact.Email, &updatedContact.Number)

	if err != nil {
		return updatedContact, err
	}

	return updatedContact, nil
}

func (r *repositoryImpl) Delete(id int64) (int64, error) {
	res, err := r.db.Exec("DELETE FROM contact_cn WHERE id_cn = $1", id)

	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return count, nil
}
