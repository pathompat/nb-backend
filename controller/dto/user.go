package dto

import "time"

type CreateUserDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"email"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	StoreName string    `json:"store"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
