package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	IsVerified bool   `gorm:"default:false"`
}
