package contact

import (
	"errors"
	"strings"
)

type Contact struct {
	Firstname string
	Lastname  string
	Email     string
	Number    string
}

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
