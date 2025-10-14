package models

type Users struct {
	ID    uint
	Name  string
	Email string
}

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
