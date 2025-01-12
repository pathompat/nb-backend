package model

import (
	"notebook-backend/types"

	"gorm.io/gorm"
)

type Production struct {
	ID          uint             `gorm:"primaryKey"`
	UserID      uint             `gorm:"not null"`
	SchoolID    uint             `gorm:"not null"`
	QuotationID uint             `gorm:"not null"`
	Remark      string           `gorm:"not null"`
	School      School           `gorm:"foreignKey:SchoolID"`
	User        User             `gorm:"references:ID"`
	Items       []ProductionItem `gorm:"foreignKey:ProductionID"`
}

type ProductionItem struct {
	gorm.Model
	ProductionID          uint           `gorm:"not null"`
	CategoryID            uint           `gorm:"not null"`
	ListQuotationConfigID types.IntArray `gorm:"type:int[];column:quotation_config_ids;default:ARRAY[]::INT[]"`
	Plate                 string         `gorm:"not null"`
	Gram                  int            `gorm:"not null"`
	Color                 string         `gorm:"not null"`
	Page                  int            `gorm:"not null"`
	Pattern               string         `gorm:"not null"`
	PrintedContent        *string        `gorm:"default:null"`
	HasReference          bool           `gorm:"not null"`
	Quantity              int            `gorm:"not null"`
	Status                string         `gorm:"not null"`
}
