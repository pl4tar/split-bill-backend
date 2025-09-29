package models

type Persons struct {
	ID   uint
	Name string
}

type PersonInput struct {
	Name string `json:"name"`
}
