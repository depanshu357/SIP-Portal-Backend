package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	CreatedAt  time.Time
	Heading    string         `gorm:"not null"`    // Ensures the Heading is not null
	Content    string         `gorm:"not null"`    // Ensures the Content is not null
	Recipients pq.StringArray `gorm:"type:text[]"` // PostgreSQL array type
	Event      uint           // PostgreSQL array type
}
