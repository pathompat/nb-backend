package dto

import (
	"time"

	"github.com/google/uuid"
)

type Production struct {
	ID     uint             `json:"id" example:"1"`            // Document id
	Remark string           `json:"remark" example:"test 123"` // Document remark
	Items  []ProductionItem `json:"items"`                     // Related items
}

type UpdateStatusItemProduction struct {
	Status string `json:"status" example:"DESIGNING"` // Document status
}

type ProductionResponse struct {
	ID              uint             `json:"id" example:"1"`                                        // Document id
	UserID          uuid.UUID        `json:"userId" example:"78705ee5-25cd-45b5-8cb1-63f1cb94e5c8"` // Owner uuid
	Username        string           `json:"userName" example:"adminTest"`                          // User name
	StoreName       string           `json:"storeName" example:"Notebook store"`                    // Store name
	SchoolName      string           `json:"schoolName" example:"School 1"`                         // School name
	SchoolAddress   string           `json:"schoolAddress" example:"33/33 Sriratch road"`           // School address
	SchoolTelephone string           `json:"schoolTelephone" example:"0812232212"`                  // School telephone
	Remark          string           `json:"remark" example:"test remark"`                          // Document remark
	Items           []ProductionItem `json:"items"`                                                 // Related items
}

type ProductionItemResponse struct {
	ID     uint   `json:"id" example:"1"`             // Document id
	ItemID int    `json:"itemId" example:"2"`         // item id of production
	Status string `json:"status" example:"DESIGNING"` // Document status
}

type ProductionItem struct {
	ProductTitle string    `json:"productTitle" example:"Cut 8"`                        // Product name
	Plate        string    `json:"plate" example:"LARGE"`                               // Plate size (LARGE, SMALL)
	Gram         int       `json:"gram" example:"40"`                                   // Notebook grams (40-150)
	Color        string    `json:"color" example:"1"`                                   // Color (1,4)
	Page         int       `json:"page" example:"40"`                                   // Page count (30-80)
	Pattern      string    `json:"pattern" example:"TABLE"`                             // Page pattern
	HasReference bool      `json:"hasReference" example:"false"`                        // Has reference
	Quantity     int       `json:"quantity" example:"1000"`                             // Product quantity
	Status       string    `json:"status" example:"PRINTING"`                           // Production status
	CreatedAt    time.Time `json:"createdAt" example:"2024-12-07T19:04:39.70268+07:00"` // Created date
	UpdatedAt    time.Time `json:"updatedAt" example:"2024-12-07T19:04:39.70268+07:00"` // Updated date
}
