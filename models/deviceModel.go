package models

import "time"

type Device struct {
	ID              uint `gorm:"primaryKey"`
	RoomID          uint
	Name            string
	LastOnlineTime  time.Time
	AttributeValues []AttributeValue
}
