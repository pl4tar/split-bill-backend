package entity

type Users struct {
	ID       uint   `json:"id" example:"1"`
	Name     string `json:"name" example:"Dima"`
	Email    string `json:"email" example:"user@exmple.com"`
	Password string `json:"password" example:"password"`
}

// type UserRegister struct {
// 	Name     string
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type UserLogin struct {
// 	Name     string
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
