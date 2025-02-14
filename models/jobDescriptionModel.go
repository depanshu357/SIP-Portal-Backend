package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type JobDescription struct {
	gorm.Model
	RecruiterID uint
	Recruiter   Recruiter `gorm:"foreignKey:RecruiterID"`
	EventID     uint
	Event       Event `gorm:"foreignKey:EventID"`
	Profile     string
	Title       string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	Location    string         `gorm:"not null"`
	Stipend     string         `gorm:"not null"`
	Eligibility pq.StringArray `gorm:"type:text[]"`
}
