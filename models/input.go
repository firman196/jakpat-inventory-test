package models

type RegisterInput struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
