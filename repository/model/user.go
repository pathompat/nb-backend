package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        int    `json:"id"`
	Username  string `json:"name" gorm:"not null"`
	StoreName string `json:"store" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	TierID    int    `json:"tierId" gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
