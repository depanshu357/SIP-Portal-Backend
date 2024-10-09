package models

import (
	"time"

	uuid "github.com/google/uuid"
	"github.com/lib/pq"
)

type Notice struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt  time.Time
	Heading    string         `gorm:"not null"`    // Ensures the Heading is not null
	Content    string         `gorm:"not null"`    // Ensures the Content is not null
	Recipients pq.StringArray `gorm:"type:text[]"` // PostgreSQL array type
	Event      uuid.UUID      `gorm:"type:uuid"`   // PostgreSQL array type
}
