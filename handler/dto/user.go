package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUser struct {
	Username  string `json:"username" binding:"required" example:"testuser123"`   // Username
	TierID    int    `json:"tierId" binding:"required" example:"1"`               // User tier (1,2,3)
	StoreName string `json:"storeName" binding:"required" example:"Example Shop"` // User's shop name
	Password  string `json:"password" binding:"required" example:"Password@1234"` // Secure password
}

type UpdateUser struct {
	Username  string `json:"username" binding:"required" example:"testuser123"`   // Username
	TierID    int    `json:"tierId" binding:"required" example:"1"`               // User tier (1,2,3)
	StoreName string `json:"storeName" binding:"required" example:"Example Shop"` // User's shop name
	Password  string `json:"password" binding:"required" example:"Password@1234"` // Secure password
}

type UserResponse struct {
	UserId    uuid.UUID `json:"id" example:"be40de0f-ba3d-44d8-9c80-023ac23e0b9a"`   // UUID generate from database
	Username  string    `json:"username" example:"testuser1"`                        // Username
	StoreName string    `json:"storeName" example:"Test Store"`                      // User's shop name
	TierID    int       `json:"tierId" example:"1"`                                  // User tier (1,2,3)
	Role      string    `json:"role" example:"CUSTOMER"`                             // User role (ADMIN, CUSTOMER)
	CreatedAt time.Time `json:"createdAt" example:"2024-12-07T19:04:39.70268+07:00"` // Created user date
	UpdatedAt time.Time `json:"updatedAt" example:"2024-12-07T19:04:39.70268+07:00"` // Latest update user date
}
