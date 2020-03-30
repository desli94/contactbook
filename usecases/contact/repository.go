package contact

type ContactRepository interface {
	Add(c ContactDTO) (ContactDTO, error)
	GetById(id int64) (ContactDTO, error)
	Update(id int64, newData ContactDTO) (ContactDTO, error)
	Delete(id int64) (int64, error)
}

type ContactDTO struct {
	ID        int64
	Firstname string
	Lastname  string
	Email     string
	Number    string
}
