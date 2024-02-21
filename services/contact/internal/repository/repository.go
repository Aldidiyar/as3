package repository

import "arch/services/contact/internal/domain"

type ContactRepository interface {
	Create(contact *domain.Contact) (*domain.Contact, error)
	Read(id int) (*domain.Contact, error)
	Update(contact *domain.Contact) (*domain.Contact, error)
	Delete(id int) error
}

type GroupRepository interface {
	Create(group *domain.Group) (*domain.Group, error)
	Read(id int) (*domain.Group, error)
	AddContact(groupID, contactID int) error
}
