package entity

type Bills struct {
	ID            uint   `json:"bill_id,string" example:"1"`
	Title         string `json:"bill_title" example:"Weekend with friends"`
	CreatedUserID uint   `json:"user_id,string" example:"1"`
}

type BillDel struct {
	ID uint `json:"bill_id,string" example:"1"`
}
