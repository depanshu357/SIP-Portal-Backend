package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title        string
	StartDate    time.Time
	IsActive     bool
	AcademicYear string
}
