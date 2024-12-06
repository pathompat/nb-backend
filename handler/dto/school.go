package dto

import (
	"time"
)

type CreateSchool struct {
	UserID    string `json:"userId" binding:"required" example:"ebf889fd-4f3c-4c15-b44b-1d37cd2ee5e4"` // UserID
	Name      string `json:"name" binding:"required" example:"school 2"`                               // Name
	Address   string `json:"address" binding:"required" example:"81 test address"`                     // Address
	Telephone string `json:"telephone" binding:"required" example:"0815231112"`                        // Telephone
}

type SchoolResponse struct {
	ID        uint      `json:"id" example:"1"`                                  // id
	Name      string    `json:"name" example:"school 1"`                         // name
	Address   string    `json:"address" example:"22/11 test address"`            // address
	Telephone string    `json:"telephone" example:"0815231112"`                  // User tier (1,2,3)
	CreatedAt time.Time `json:"createdAt" example:"2024-12-02T00:26:21.087061Z"` // Created user date
	UpdatedAt time.Time `json:"updatedAt" example:"2024-12-02T00:26:21.087061Z"` // Latest update user date
}
