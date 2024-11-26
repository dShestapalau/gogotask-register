package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	ID         uuid.UUID
	LastName   string
	FirstName  string
	MiddleName string
	Locale     string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
