package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUser struct {
	Username  string `json:"username" binding:"required"`
	TierID    int    `json:"tierId" binding:"required"`
	StoreName string `json:"storeName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUser struct {
	Username  string `json:"username" binding:"required"`
	TierID    int    `json:"tierId" binding:"required"`
	StoreName string `json:"storeName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserResponse struct {
	UserId    uuid.UUID `json:"id" example:"be40de0f-ba3d-44d8-9c80-023ac23e0b9a"`
	Username  string    `json:"username" example:"testuser1"`
	StoreName string    `json:"storeName" example:"Test Store"`
	TierID    int       `json:"tierId" example:"1"`
	Role      string    `json:"role" example:"CUSTOMER"`
	CreatedAt time.Time `json:"createdAt" example:"2024-12-02T00:26:21.087061Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-12-02T00:26:21.087061Z"`
}
