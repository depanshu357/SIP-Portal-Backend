package models

import "gorm.io/gorm"

type Applicant struct {
	gorm.Model
	FileID           uint
	File             File `gorm:"foreignKey:FileID"`
	JobDescriptionID uint
	JobDescription   JobDescription `gorm:"foreignKey:JobDescriptionID"`
	StudentID        uint
	Student          Student `gorm:"foreignKey:StudentID"`
}
