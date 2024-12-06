package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    uuid.UUID `gorm:"column:uuid;type:uuid;default:uuid_generate_v4()"`
	Username  string    `gorm:"not null"`
	StoreName string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null;default:CUSTOMER"`
	TierID    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
