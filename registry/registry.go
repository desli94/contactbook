package registry

import (
	"database/sql"

	ac "github.com/desli94/contactbook/adapters/contact"
	uc "github.com/desli94/contactbook/usecases/contact"
)

type registry struct {
	db *sql.DB
}

// Registry interface specifies the controllers that should be created
type Registry interface {
	NewContactController() ac.Controller
}

// NewRegistry create the registry that holds all the controllers
func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewContactController() ac.Controller {
	contactRepo := ac.NewRepo(r.db)
	contactInteractor := uc.NewContactInteractor(contactRepo)
	return ac.NewContactController(contactInteractor)
}
