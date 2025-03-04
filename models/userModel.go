package models

import (
	"time"

	"encoding/json"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreatedAt         time.Time
	Email             string          `gorm:"unique;not null"`
	Password          string          `gorm:"not null"`
	Role              string          `gorm:"default:student"`
	IsProfileVerified bool            `gorm:"default:false"`
	HasAdminAccess    bool            `gorm:"default:false"`
	VerifiedForEvents pq.Int64Array   `gorm:"type:integer[];default:'{}'"`
	FrozenForEvents   pq.Int64Array   `gorm:"type:integer[];default:'{}'"`
	ReasonForFreeze   json.RawMessage `gorm:"type:jsonb;default:'{}'::jsonb"`
}
