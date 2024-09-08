package models

import "time"

type Otp struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique;not null"`
	Otp          string `gorm:"not null"`
	DeletionTime time.Time
}
