package dto

import (
	"time"

	"github.com/google/uuid"
)

type QuotationFilter struct {
	IncludeProduction bool `form:"includeProduction"`
}

type CreateQuotation struct {
	UserID        uuid.UUID       `json:"userId" binding:"required,uuid4" example:"78705ee5-25cd-45b5-8cb1-63f1cb94e5c8"` // Owner uuid
	SchoolID      uint            `json:"schoolId" binding:"required,gt=0" example:"2"`                                   // School id
	AppointmentAt *time.Time      `json:"appointmentAt" binding:"-" example:"2024-12-00:00:00.0000+07:00"`                // Appointment date (null is now)
	DueDateAt     time.Time       `json:"dueDateAt" binding:"required" example:"2024-12-06"`                              // Last due date
	Items         []QuotationItem `json:"items" binding:"required,dive"`                                                  // Quotation product list
	Remark        string          `json:"remark" example:"remark test"`                                                   // Any remark
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
	Remark          string              `json:"remark" example:"test remark"`                          // Document remark
	Production      *ProductionResponse `json:"production,omitempty"`                                  // Production related
	Items           []QuotationItem     `json:"items"`                                                 // Quotation product list
	CreatedAt       time.Time           `json:"createdAt" example:"2024-12-07T19:04:39.70268+07:00"`   // Created date
	UpdatedAt       time.Time           `json:"updatedAt" example:"2024-12-07T19:04:39.70268+07:00"`   // Latest update date
}

type QuotationItem struct {
	ProductTitle string  `json:"productTitle" binding:"required" example:"Cut 8"`      // Product name
	Plate        string  `json:"plate" binding:"-" example:"LARGE"`                    // Plate size (LARGE, SMALL)
	Gram         int     `json:"gram" binding,gte=5:"required" example:"40"`           // Notebook grams (40-150)
	Color        string  `json:"color" binding:"required" example:"1"`                 // Color (1,4)
	Page         int     `json:"page" binding,gte=10:"required" example:"40"`          // Page count (30-80)
	Pattern      string  `json:"pattern" binding:"required,uppercase" example:"TABLE"` // Page pattern
	HasReference *bool   `json:"hasReference" binding:"required" example:"false"`      // Has reference
	Quantity     int     `json:"quantity" binding:"required,gte=1" example:"1000"`     // Product quantity
	Price        float32 `json:"price" binding:"required,gt=0" example:"5.5"`          // Product price
}
