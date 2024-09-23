package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string `gorm:"unique;not null"`
	RollNo     string `gorm:"unique"`
	IsVerified bool   `gorm:"default:false"`
	Department string
	Branch     string
}
