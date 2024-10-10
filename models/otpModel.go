package models

import (
	"time"
)

type Otp struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	Email        string `gorm:"unique;not null"`
	Otp          string `gorm:"not null"`
	DeletionTime time.Time
}
