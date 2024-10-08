package models

import uuid "github.com/google/uuid"

type Admin struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid"`
	FirstName  string
	LastName   string
	Email      string `gorm:"unique;not null"`
	IsVerified bool   `gorm:"default:false"`
}
