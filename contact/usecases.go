package contact

type repository interface {
	add(c ContactDTO) (ContactDTO, error)
}

type useCase struct {
	Repository repository
}

type ContactDTO struct {
	Firstname string
	Lastname  string
	Email     string
	Number    string
}

func NewUseCase(rep repository) useCase {
	return useCase{rep}
}

func (us useCase) Add(c ContactDTO) (ContactDTO, error) {
	cont, err := MakeNewContact(Contact{c.Firstname, c.Lastname, c.Email, c.Number})

	if err != nil {
		return ContactDTO{}, err
	}
	contDTO, err := us.Repository.add(ContactDTO{cont.Firstname, cont.Lastname, cont.Email, cont.Number})

	if err != nil {
		return ContactDTO{}, err
	}

	return contDTO, nil
}
