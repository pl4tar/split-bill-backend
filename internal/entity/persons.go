package entity

type Persons struct {
	ID      uint   `json:"id,string" example:"1"`
	Name    string `json:"name" example:"Danil"`
	OwnerID uint   `json:"owner_id,string" example:"1"`
}
type PersonsOutput struct {
	ID   uint   `json:"id,string" example:"1"`
	Name string `json:"name" example:"Danil"`
}
