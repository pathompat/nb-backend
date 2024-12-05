package dto

import (
	"time"
)

type SchoolResponse struct {
	Name      string    `json:"name" example:"school 1"`                         // name
	Address   string    `json:"address" example:"22/11 test address"`            // address
	Telephone string    `json:"telephone" example:"0815231112"`                  // User tier (1,2,3)
	CreatedAt time.Time `json:"createdAt" example:"2024-12-02T00:26:21.087061Z"` // Created user date
	UpdatedAt time.Time `json:"updatedAt" example:"2024-12-02T00:26:21.087061Z"` // Latest update user date
}
