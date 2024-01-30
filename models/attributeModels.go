package models

import "time"

// Attribute model
type Attribute struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	AttributeValues []AttributeValue
}

// AttributeValue model
type AttributeValue struct {
	ID          uint `gorm:"primaryKey"`
	DeviceID    uint
	AttributeID uint
	Value       string
	LastUpdate  time.Time
}
