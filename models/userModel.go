package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email             string        `gorm:"unique;not null"`
	Password          string        `gorm:"not null"`
	IsVerified        bool          `gorm:"default:false"`
	Role              string        `gorm:"default:student"`
	VerifiedForEvents pq.Int32Array `gorm:"type:integer[];default:'{}'"`
}
