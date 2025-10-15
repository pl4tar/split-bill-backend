package entity

type Users struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"Dima"`
	Email string `json:"email" example:"exmple@mail.com"`
}

type UserRegister struct {
	Name     string
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Name     string
	Email    string `json:"email"`
	Password string `json:"password"`
}
