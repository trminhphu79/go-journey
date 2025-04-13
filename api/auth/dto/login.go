package dto

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRessponseDTO struct {
	AccessToken  string `json:"accessToken" `
	RefreshToken string `json:"refreshToken" `
}
