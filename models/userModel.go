package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreatedAt         time.Time
	Email             string        `gorm:"unique;not null"`
	Password          string        `gorm:"not null"`
	IsVerified        bool          `gorm:"default:false"`
	Role              string        `gorm:"default:student"`
	VerifiedForEvents pq.Int64Array `gorm:"type:integer[];default:'{}'"`
}
