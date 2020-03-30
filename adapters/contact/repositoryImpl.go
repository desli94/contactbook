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
	rows := r.db.QueryRow("INSERT INTO contact_cn (firstname_cn, lastname_cn, email_cn, number_cn) VALUES ($1, $2, $3, $4) RETURNING *", c.Firstname, c.Lastname, c.Email, c.Number)

	contact := uc.ContactDTO{}
	err := rows.Scan(&contact.ID, &contact.Firstname, &contact.Lastname, &contact.Email, &contact.Number)

	if err != nil {
		fmt.Println("yep")
		return uc.ContactDTO{}, err
	}

	return contact, nil
}

func (r *repositoryImpl) GetById(id int64) (uc.ContactDTO, error) {
	rows, err := r.db.Query("SELECT * FROM contact_cn WHERE id_cn=$1", id)

	if err != nil {
		return uc.ContactDTO{}, err
	}

	contact := uc.ContactDTO{}
	for rows.Next() {
		err = rows.Scan(&contact.ID, &contact.Firstname, &contact.Lastname, &contact.Email, &contact.Number)

		if err != nil {
			return uc.ContactDTO{}, err
		}
	}

	return contact, nil
}

func (r *repositoryImpl) Update(id int64, newData uc.ContactDTO) (uc.ContactDTO, error) {
	rows := r.db.QueryRow("UPDATE contact_cn SET firstname_cn = $1, lastname_cn = $2, email_cn = $3, number_cn = $4 WHERE id_cn = $5 RETURNING *",
		newData.Firstname, newData.Lastname, newData.Email, newData.Number, id)

	updatedContact := uc.ContactDTO{}
	err := rows.Scan(&updatedContact.ID, &updatedContact.Firstname, &updatedContact.Lastname, &updatedContact.Email, &updatedContact.Number)

	if err != nil {
		return uc.ContactDTO{}, err
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
