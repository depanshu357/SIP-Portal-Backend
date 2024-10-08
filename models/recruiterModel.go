package models

import (
	"gorm.io/gorm"
)

type Recruiter struct {
	gorm.Model
	FirstName         string
	LastName          string
	Email             string `gorm:"unique;not null"`
	Company           string
	IsVerified        bool `gorm:"default:false"`
	IsProfileVerified bool `gorm:"default:false"`
}
