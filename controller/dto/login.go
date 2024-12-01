package dto

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseWithToken struct {
	Token     string `json:"token"`
	ExpiredIn int    `json:"expiredIn"`
}
