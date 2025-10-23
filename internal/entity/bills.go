package entity

type Bills struct {
	ID            uint   `json:"bill_id" example:"1"`
	Title         string `json:"bill_title" example:"Weekend with friends"`
	CreatedUserID uint   `json:"user_id" example:"1"`
}

type BillsInput struct {
	Title         string `json:"bill_title" example:"dfsfsd"`
	CreatedUserID string `json:"user_id" example:"1"`
}

// func New(title string, creatorID uint) *bills {
// 	return &bills{Title: title, CreatedUserID: creatorID}
// }
