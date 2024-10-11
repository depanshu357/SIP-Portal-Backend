package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID       uint
	User         User `gorm:"foreignKey:UserID"`
	Name         string
	IsVerified   bool `gorm:"default:false"`
	Event        string
	AcademicYear string
	Path         string
	Category     string
}
