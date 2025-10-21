package entity

type payments struct {
	id         uint
	checkID    uint
	fromPerson *Persons
	toPerson   *Persons
	amount     float64
	status     bool
}

//type PersonInput struct {
//
//}

func New(ID uint, CheckID uint, FromPerson *Persons, ToPerson *Persons, Amount float64, Status bool) *payments {
	return &payments{id: ID, checkID: CheckID, fromPerson: FromPerson, toPerson: ToPerson, amount: Amount, status: Status}
}
