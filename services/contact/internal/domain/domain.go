package domain

import (
	"fmt"
)

// Контакт представляет собой запись о контакте.
type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
}

// FullName возвращает полное имя контакта.
func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

// SetFullName устанавливает полное имя контакта из составных частей.
func (c *Contact) SetFullName(firstName, lastName string) {
	c.FirstName = firstName
	c.LastName = lastName
}

// Группа представляет собой группу контактов.
type Group struct {
	ID   int
	Name string
}

// NewContact создает новый экземпляр контакта.
func NewContact(id int, firstName, lastName, phoneNumber string) *Contact {
	return &Contact{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}
}

// NewGroup создает новый экземпляр группы контактов.
func NewGroup(id int, name string) *Group {
	return &Group{
		ID:   id,
		Name: name,
	}
}
