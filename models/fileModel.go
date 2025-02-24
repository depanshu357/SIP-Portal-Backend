package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
	Name       string
	IsVerified bool `gorm:"default:false"`
	EventID    uint
	Event      Event `gorm:"foreignKey:EventID"`
	Path       string
	Category   string
}
