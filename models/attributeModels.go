package models

import (
	"time"

	"github.com/google/uuid"
)

// Attribute model
type Attribute struct {
	ID              uuid.UUID        `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	Name            string           `gorm:"not null;unique;"`
	AttributeValues []AttributeValue `gorm:"foreignKey:AttributeName;references:Name;"`
	CreatedAt       time.Time        `gorm:"autoCreateTime"`
}

// AttributeValue model
type AttributeValue struct {
	ID            uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	Value         string
	DeviceID      uuid.UUID `gorm:"not null;"`
	AttributeName string    `gorm:"not null;"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
