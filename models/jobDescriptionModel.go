package models

import "gorm.io/gorm"

type JobDescription struct {
	gorm.Model
	RecruiterID uint
	Recruiter   Recruiter `gorm:"foreignKey:RecruiterID"`
	EventID     uint
	Event       Event  `gorm:"foreignKey:EventID"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Location    string `gorm:"not null"`
	Stipend     string `gorm:"not null"`
}
