package models

import uuid "github.com/google/uuid"

type Recruiter struct {
	ID                uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name              string
	Email             string `gorm:"unique;not null"`
	Company           string
	IsVerified        bool `gorm:"default:false"`
	IsProfileVerified bool `gorm:"default:false"`
}
