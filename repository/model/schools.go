package model

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	UserId    int       `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Address   string    `gorm:"not null"`
	Telephone string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
