package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	LastName   string
	FirstName  string
	MiddleName string
	Locale     string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (UserProfile) TableName() string {
	return "user_profile"
}
