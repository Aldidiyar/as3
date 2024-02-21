package useCase

type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
}

type Group struct {
	ID   int
	Name string
}

type ContactUseCase interface {
	CreateContact(firstName, lastName, phoneNumber string) (*Contact, error)
	ReadContact(id int) (*Contact, error)
	UpdateContact(id int, firstName, lastName, phoneNumber string) (*Contact, error)
	DeleteContact(id int) error
}

type GroupUseCase interface {
	CreateGroup(name string) (*Group, error)
	ReadGroup(id int) (*Group, error)
	AddContactToGroup(groupID, contactID int) error
}
