package models

type Bills struct {
	id           uint
	title        string
	CreateUserID uint
	People       []Persons
	Products     []Products
	// ToDo Assignments для подсчета кто кому должен
}
