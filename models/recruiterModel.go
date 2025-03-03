package models

import (
	"gorm.io/gorm"
)

type Recruiter struct {
	gorm.Model
	UserID           uint
	User             User `gorm:"foreignKey:UserID"`
	Name             string
	Email            string `gorm:"unique;not null"`
	Company          string
	ContactNumber    string
	NatureOfBusiness string
	AdditionalInfo   string
}
