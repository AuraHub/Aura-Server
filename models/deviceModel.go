package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID              uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	RoomID          uuid.UUID
	Name            string
	Online          bool
	LastOnline      time.Time
	AttributeValues []AttributeValue
	Configured      bool `gorm:"default:false"`
}
