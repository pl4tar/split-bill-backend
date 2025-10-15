package entity

type Persons struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"Danil"`
}

type PersonInput struct {
	Name string `json:"name" example:"Danil"`
}
