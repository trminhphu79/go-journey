package dto

type RegistrationDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullName" bining:"required"`
}
