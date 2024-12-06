package model

import (
	"time"

	"gorm.io/gorm"
)

type Quotation struct {
	gorm.Model
	UserID          uint            `gorm:"not null"`
	SchoolID        uint            `gorm:"not null"`
	StoreName       string          `gorm:"not null"`
	SchoolName      string          `gorm:"not null"`
	SchoolAddress   string          `gorm:"not null"`
	SchoolTelephone string          `gorm:"not null"`
	AppointmentAt   *time.Time      `gorm:"not null"`
	DueDateAt       time.Time       `gorm:"column:duedate_at, not null"`
	Status          string          `gorm:"not null"`
	Remark          string          `gorm:"not null"`
	Production      *Production     `gorm:"foreignKey:QuotationID"`
	Items           []QuotationItem `gorm:"foreignKey:QuotationID"`
}

type QuotationItem struct {
	ID           uint   `gorm:"primaryKey"`
	QuotationID  uint   `gorm:"not null"`
	ProductTitle string `gorm:"not null"`
	Plate        string
	Gram         int     `gorm:"not null"`
	Color        string  `gorm:"not null"`
	Page         int     `gorm:"not null"`
	Pattern      string  `gorm:"not null"`
	HasReference bool    `gorm:"not null"`
	Quantity     int     `gorm:"not null"`
	Price        float32 `gorm:"not null"`
}
