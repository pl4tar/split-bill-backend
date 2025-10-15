package entity

type Bills struct {
	ID            string `json:"bill_id" example:"1"`
	Title         string `json:"bill_title" example:"Weekend with friends"`
	CreatedUserID uint   `json:"user_id" example:"1"`
}
