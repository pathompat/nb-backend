package model

import (
	"gorm.io/gorm"
)

type PriceReference struct {
	gorm.Model
	TierID       int     `gorm:"not null"`
	Category     string  `gorm:"not null"`
	Plate        string  `gorm:"not null"`
	Gram         int     `gorm:"not null"`
	Color        string  `gorm:"not null"`
	Page         int     `gorm:"not null"`
	Pattern      string  `gorm:"not null"`
	HasReference bool    `gorm:"not null"`
	Price        float64 `gorm:"not null"`
}
