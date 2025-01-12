package model

import (
	"notebook-backend/types"
	"time"

	"gorm.io/gorm"
)

type Quotation struct {
	gorm.Model
	UserID            uint            `gorm:"not null"`
	User              User            `gorm:"references:ID"`
	SchoolID          uint            `gorm:"not null"`
	StoreName         string          `gorm:"not null"`
	SchoolName        string          `gorm:"not null"`
	SchoolAddress     *string         `gorm:"default:null"`
	SchoolTelephone   *string         `gorm:"default:null"`
	SchoolContactName *string         `gorm:"default:null"`
	AppointmentAt     *time.Time      `gorm:"default:null"`
	DueDateAt         time.Time       `gorm:"not null;column:duedate_at"`
	Status            string          `gorm:"not null;default:'REVIEWING'"`
	Remark            string          `gorm:"default:null"`
	Production        *Production     `gorm:"foreignKey:QuotationID"`
	Items             []QuotationItem `gorm:"foreignKey:QuotationID"`
}

type QuotationItem struct {
	ID                uint           `gorm:"primaryKey"`
	QuotationID       uint           `gorm:"not null"`
	CategoryID        uint           `gorm:"not null"`
	QuotationConfigID types.IntArray `gorm:"type:int[];column:quotation_config_ids;default:ARRAY[]::INT[]"`
	Category          Category       `gorm:"references:ID"`
	Plate             *string        `gorm:"default:null"`
	Gram              int            `gorm:"not null"`
	Color             string         `gorm:"not null"`
	Page              int            `gorm:"not null"`
	Pattern           string         `gorm:"not null"`
	PrintedContent    *string        `gorm:"default:null"`
	HasReference      bool           `gorm:"not null"`
	Quantity          int            `gorm:"not null"`
	Price             float32        `gorm:"not null"`
	Charge            float32        `gorm:"not null;default:0"`
}

type QuotationConfig struct {
	gorm.Model
	CategoryID       *uint          `gorm:"default:null"`
	Category         Category       `gorm:"references:ID"`
	TierIDList       types.IntArray `gorm:"type:smallint[];column:tier_ids;default:ARRAY[]::SMALLINT[]"`
	Color            *string        `gorm:"default:NULL"`
	Key              string         `gorm:"not null;unique"`
	Value            float32        `gorm:"not null"`
	Description      *string        `gorm:"default:null"`
	Label            *string        `gorm:"default:null"`
	Unit             string         `gorm:"not null"`
	Level            string         `gorm:"not null"`
	Comparator       *string        `gorm:"default:null"`
	CompareValue     int            `gorm:"default:null"`
	HasFixedCharge   bool           `gorm:"default:FALSE"`
	FixedChargePrice float32        `gorm:"default:0"`
	Type             *string        `gorm:"default:null"`
}

type StatusCount struct {
	Status string
	Count  int
}
