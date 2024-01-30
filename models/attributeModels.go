package models

import (
	"time"

	"github.com/google/uuid"
)

// Attribute model
type Attribute struct {
	ID              uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	Name            string
	AttributeValues []AttributeValue
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}

// AttributeValue model
type AttributeValue struct {
	ID          uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	Value       string
	DeviceID    uuid.UUID
	AttributeID uuid.UUID
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
