package contact

import (
	"errors"
	"strings"
)

// Contact entity
type Contact struct {
	ID        int64
	Firstname string
	Lastname  string
	Email     string
	Number    string
}

// MakeNewContact validates and creates a contact entity.
func MakeNewContact(c Contact) (Contact, error) {
	errs := []string{}
	if c.Firstname == "" {
		errs = append(errs, "firstname empty")
	}
	if c.Lastname == "" {
		errs = append(errs, "lastname empty")
	}
	if c.Email == "" {
		errs = append(errs, "email empty")
	}
	if c.Number == "" {
		errs = append(errs, "number empty")
	}
	if len(errs) > 0 {
		return Contact{}, errors.New("Err: " + strings.Join(errs, "\n"))
	}

	return c, nil
}
