package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName  string `gorm:"not null"`
	LastName   string
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	RollNo     string `gorm:"unique"`
	IsVerified bool   `gorm:"default:false"`
	Department string
	Branch     string
}
