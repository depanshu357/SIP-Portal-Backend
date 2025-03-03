package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
	Name   string
	Email  string `gorm:"unique;not null"`
}
