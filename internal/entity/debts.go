package entity

type Debt struct {
	FromPersonID   uint    `json:"from_person_id"`
	FromPersonName string  `json:"from_person_name"`
	ToPersonID     uint    `json:"to_person_id"`
	ToPersonName   string  `json:"to_person_name"`
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
}

type DebtCalculation struct {
	BillID uint   `json:"bill_id"`
	Debts  []Debt `json:"debts"`
}
