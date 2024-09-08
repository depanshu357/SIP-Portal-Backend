package models

import (
	"gorm.io/gorm"
)

type Recruiter struct {
	gorm.Model
	FirstName  string `gorm:"not null"`
	LastName   string
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Company    string
	IsVerified bool `gorm:"default:false"`
}
