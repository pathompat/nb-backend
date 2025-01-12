package model

import (
	"notebook-backend/types"

	"gorm.io/gorm"
)

type PriceReference struct {
	gorm.Model
	TierID     int            `gorm:"not null"`
	CategoryID uint           `gorm:"not null"`
	Category   Category       `gorm:"references:ID"`
	Gram       int            `gorm:"not null"`
	Color      *string        `gorm:"default:null"`
	Page       int            `gorm:"not null"`
	Pattern    types.StrArray `gorm:"type:TEXT[];default:ARRAY[]::TEXT[]"`
	Price      float64        `gorm:"not null"`
}
