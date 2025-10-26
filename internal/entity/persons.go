package entity

type Persons struct {
	ID     uint   `json:"id,string" example:"1"`
	Name   string `json:"name" example:"Danil"`
	BillID uint   `json:"bill_id,string" example:"1"`
}
type PersonsOutput struct {
	ID   uint   `json:"id,string" example:"1"`
	Name string `json:"name" example:"Danil"`
}
