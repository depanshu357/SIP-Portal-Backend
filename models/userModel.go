package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreatedAt         time.Time
	Email             string         `gorm:"unique;not null"`
	Password          string         `gorm:"not null"`
	Role              string         `gorm:"default:student"`
	IsProfileVerified bool           `gorm:"default:false"`
	HasAdminAccess    bool           `gorm:"default:false"`
	VerifiedForEvents pq.Int64Array  `gorm:"type:integer[];default:'{}'"`
	FrozenForEvents   pq.Int64Array  `gorm:"type:integer[];default:'{}'"`
	ReasonForFreeze   pq.StringArray `gorm:"type:text[];default:'{}'"`
}
