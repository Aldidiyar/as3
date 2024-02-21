package domain

import (
	"fmt"
)

type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func (c *Contact) SetFullName(firstName, lastName string) {
	c.FirstName = firstName
	c.LastName = lastName
}

type Group struct {
	ID   int
	Name string
}

func NewContact(id int, firstName, lastName, phoneNumber string) *Contact {
	return &Contact{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}
}

func NewGroup(id int, name string) *Group {
	return &Group{
		ID:   id,
		Name: name,
	}
}
