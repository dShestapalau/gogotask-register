package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserCredentials struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time

	UserProfile UserProfile `gorm:"foreignKey:UserID"`
}
