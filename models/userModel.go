package models

import (
	"time"

	uuid "github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID                uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt         time.Time
	Email             string         `gorm:"unique;not null"`
	Password          string         `gorm:"not null"`
	IsVerified        bool           `gorm:"default:false"`
	Role              string         `gorm:"default:student"`
	VerifiedForEvents pq.StringArray `gorm:"type:uuid[];default:'{}'"`
}
