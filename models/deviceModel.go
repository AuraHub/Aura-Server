package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID              uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	DeviceId        string    `gorm:"not null;unique;"`
	Name            string
	RoomID          uuid.UUID `gorm:"foreignKey:ID;default:null;"`
	Online          bool      `gorm:"default:true;not null;"`
	LastOnline      time.Time `gorm:"autoCreateTime;not null;"`
	AttributeValues []AttributeValue
	Configured      bool `gorm:"default:false;not null;"`
}
