package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title        string
	StartDate    time.Time
	IsActive     bool
	AcademicYear string
}
