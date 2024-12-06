package dto

import (
	"time"

	"github.com/google/uuid"
)

type QuotationFilter struct {
	IncludeProduction bool `form:"includeProduction"`
}

type QuotationResponse struct {
	ID              uint                `json:"id" example:"1"`                                        // Document id
	UserID          uuid.UUID           `json:"userId" example:"78705ee5-25cd-45b5-8cb1-63f1cb94e5c8"` // Owner uuid
	StoreName       string              `json:"storeName" example:"Notebook store"`                    // Store name
	SchoolName      string              `json:"schoolName" example:"School 1"`                         // School name
	SchoolAddress   string              `json:"schoolAddress" example:"33/33 Sriratch road"`           // School address
	SchoolTelephone string              `json:"schoolTelephone" example:"0812232212"`                  // School telephone
	AppointmentAt   *time.Time          `json:"appointmentAt" example:"2024-12-02"`                    // Appointment date (null is now)
	DueDateAt       time.Time           `json:"dueDateAt" example:"2024-12-02"`                        // Last due date
	Status          string              `json:"status" example:"REVIEWING"`                            // Document status (REVIEWING, APPROVED, CANCELED)
	Production      *ProductionResponse `json:"production,omitempty"`                                  // Production related
	Items           []QuotationItem     `json:"items"`                                                 // Quotation product list
	CreatedAt       time.Time           `json:"createdAt" example:"2024-12-02T00:26:21.087061Z"`       // Created date
	UpdatedAt       time.Time           `json:"updatedAt" example:"2024-12-02T00:26:21.087061Z"`       // Latest update date
}

type QuotationItem struct {
	ProductTitle string  `json:"productTitle" example:"Cut 8"` // Product name
	Plate        string  `json:"plate" example:"LARGE"`        // Plate size (LARGE, SMALL)
	Gram         int     `json:"gram" example:"40"`            // Notebook grams (40-150)
	Color        string  `json:"color" example:"1"`            // Color (1,4)
	Page         int     `json:"page" example:"40"`            // Page count (30-80)
	Pattern      string  `json:"pattern" example:"TABLE"`      // Page pattern
	HasReference bool    `json:"hasReference" example:"false"` // Has reference
	Quantity     int     `json:"quantity" example:"1000"`      // Product quantity
	Price        float32 `json:"price" example:"5.5"`          // Product price
}
