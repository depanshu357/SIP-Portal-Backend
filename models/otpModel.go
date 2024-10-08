package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Otp struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	Email        string `gorm:"unique;not null"`
	Otp          string `gorm:"not null"`
	DeletionTime time.Time
}
