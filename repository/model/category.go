package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Key         string `gorm:"not null;unique"`
	NameTH      string `gorm:"not null"`
	NameEN      string `gorm:"not null"`
	ListGram    string `gorm:"not null"`
	ListPage    string `gorm:"not null"`
	ListPattern string `gorm:"not null"`
}
