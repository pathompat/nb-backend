package dto

import (
	"time"
)

type ProductionResponse struct {
	ID     uint             `json:"id" example:"1"`            // Document id
	Remark string           `json:"remark" example:"test 123"` // Document remark
	Items  []ProductionItem `json:"items"`                     // Related items
}

type ProductionItem struct {
	ProductTitle string    `json:"productTitle" example:"Cut 8"`                    // Product name
	Plate        string    `json:"plate" example:"LARGE"`                           // Plate size (LARGE, SMALL)
	Gram         int       `json:"gram" example:"40"`                               // Notebook grams (40-150)
	Color        string    `json:"color" example:"1"`                               // Color (1,4)
	Page         int       `json:"page" example:"40"`                               // Page count (30-80)
	Pattern      string    `json:"pattern" example:"TABLE"`                         // Page pattern
	HasReference bool      `json:"hasReference" example:"false"`                    // Has reference
	Quantity     int       `json:"quantity" example:"1000"`                         // Product quantity
	Status       string    `json:"status" example:"PRINTING"`                       // Production status
	CreatedAt    time.Time `json:"createdAt" example:"2024-12-02T00:26:21.087061Z"` // Created date
	UpdatedAt    time.Time `json:"updatedAt" example:"2024-12-02T00:26:21.087061Z"` // Updated date
}
