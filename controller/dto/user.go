package dto

import "time"

type CreateUserDTO struct {
	Username  string  `json:"username" binding:"required"`
	TierID    int     `json:"tierId" binding:"required"`
	StoreName string  `json:"store" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}

type UpdateUserDTO struct {
	Username  string  `json:"username" binding:"required"`
	TierID    int     `json:"tierId" binding:"required"`
	StoreName string  `json:"store" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}

type DeleteUserDTO struct {
	ID        int      `json:"id"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	StoreName string    `json:"storeName"`
	TierID    int    	`json:"tierId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
