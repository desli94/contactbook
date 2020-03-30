package contact

type ContactRepository interface {
	Add(c ContactDTO) (ContactDTO, error)
}
