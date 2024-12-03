package dto

type Login struct {
	Username string `json:"username" binding:"required,lowercase,alphanum" example:"testuser123"` // Username
	Password string `json:"password" binding:"required,min=8" example:"Password@123"`             // Password
}

type ResponseWithToken struct {
	Token     string `json:"token" example:"token123"` // JWT Token
	ExpiredIn int    `json:"expiredIn" example:"3600"` // Token expired in (second)
}
