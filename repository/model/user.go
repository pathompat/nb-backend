package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"name" gorm:"not null"`
	StoreName string `json:"store" gorm:"not null"`
	TierID    int    `json:"tier" gorm:"not null"`
}
