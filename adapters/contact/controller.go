package contact

import (
	ec "github.com/desli94/contactbook/entities/contact"
	uc "github.com/desli94/contactbook/usecases/contact"
)

// type contact struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// 	Email     string `json:"email"`
// 	Number    string `json:"number"`
// }

type controller struct {
	Interactor uc.Interactor
}

// Controller interface specifies the methods to be called by infrastructure
type Controller interface {
	Add(firstname string, lastname string, email string, number string) (ec.Contact, error)
	GetAll() ([]ec.Contact, error)
}

var contactRepo uc.Repository

// NewContactController creates the controller with the injected dependecies
func NewContactController(ci uc.Interactor) Controller {
	return &controller{ci}
}

// func NewContactController(db *sql.DB) {
// 	contactRepo = NewRepo(db)
// }

func (c *controller) Add(firstname string, lastname string, email string, number string) (ec.Contact, error) {
	var contact ec.Contact

	contact = ec.Contact{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Number:    number,
	}

	addedContact, err := c.Interactor.Add(contact)

	if err != nil {
		return addedContact, err
	}

	return addedContact, nil
	// json.NewEncoder(w).Encode(addedCOntact)
}

func (c *controller) GetAll() ([]ec.Contact, error) {
	contacts, err := c.Interactor.GetAll()

	if err != nil {
		return contacts, err
	}

	return contacts, nil
}

// func (c *controll) Add(w http.ResponseWriter, r *http.Request) {
// 	var c ec.Contact
// 	json.NewDecoder(r.Body).Decode(&c)

// 	ci := uc.NewContactInteractor(contactRepo)
// 	addedCOntact, err := ci.Add(c)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	json.NewEncoder(w).Encode(addedCOntact)
// }
