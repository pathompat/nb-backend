package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Username  string `json:"username" binding:"required"`
	TierID    int    `json:"tierId" binding:"required"`
	StoreName string `json:"storeName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserDTO struct {
	Username  string `json:"username" binding:"required"`
	TierID    int    `json:"tierId" binding:"required"`
	StoreName string `json:"storeName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type User struct {
	UserId    uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	StoreName string    `json:"storeName"`
	TierID    int       `json:"tierId" gorm:"not null"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
