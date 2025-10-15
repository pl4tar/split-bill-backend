package entity

type Payments struct {
	ID         uint
	CheckID    uint
	FromPerson Persons
	ToPerson   Persons
	Amount     float64
	Status     bool
}

//type PersonInput struct {
//
//}
