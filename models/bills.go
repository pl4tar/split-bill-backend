package models

import "time"

type Bills struct {
	id           uint
	title        string
	CreateUserID uint
	CreateAt     time.Time
	People       []Persons
	Products     []Products
	// ToDo Assignments для подсчета кто кому должен
}
