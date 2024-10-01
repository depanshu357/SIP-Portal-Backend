package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	Heading    string         `gorm:"not null"`    // Ensures the Heading is not null
	Content    string         `gorm:"not null"`    // Ensures the Content is not null
	Recipients pq.StringArray `gorm:"type:text[]"` // PostgreSQL array type
}
