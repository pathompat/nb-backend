package model

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	Name        string    `gorm:"not null"`
	ContactName string    `gorm:"-"`
	Address     string    `gorm:"-"`
	Telephone   string    `gorm:"-"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
