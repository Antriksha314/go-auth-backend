package dto

type RegisterDTO struct {
	FirstName string `json:"firstName" binded:"required"`
	LastName  string `json:"lastName" binded:"required"`
	Email     string `json:"email" binded:"required"`
	Password  string `json:"password" binded:"required"`
}
