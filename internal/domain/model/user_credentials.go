package model

import (
	"time"

	"github.com/google/uuid"
)

type UserCredentials struct {
	UserId    uuid.UUID
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
