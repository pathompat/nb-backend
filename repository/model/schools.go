package model

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	Name        string    `gorm:"not null"`
	ContactName *string   `gorm:"default:null"`
	Address     *string   `gorm:"default:null"`
	Telephone   *string   `gorm:"default:null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
