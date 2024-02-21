package useCase

// Contact представляет запись о контакте.
type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
}

// Group представляет группу контактов.
type Group struct {
	ID   int
	Name string
}

// ContactUseCase представляет интерфейс для Use Case контакта.
type ContactUseCase interface {
	CreateContact(firstName, lastName, phoneNumber string) (*Contact, error)
	ReadContact(id int) (*Contact, error)
	UpdateContact(id int, firstName, lastName, phoneNumber string) (*Contact, error)
	DeleteContact(id int) error
}

// GroupUseCase представляет интерфейс для Use Case группы контактов.
type GroupUseCase interface {
	CreateGroup(name string) (*Group, error)
	ReadGroup(id int) (*Group, error)
	AddContactToGroup(groupID, contactID int) error
}
