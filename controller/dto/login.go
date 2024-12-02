package dto

type Login struct {
	Username string `json:"username" binding:"required,lowercase,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
}

type ResponseWithToken struct {
	Token     string `json:"token"`
	ExpiredIn int    `json:"expiredIn"`
}
