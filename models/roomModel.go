package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	Name      string    `gorm:"not null;"`
	Devices   []Device
	Users     []*User   `gorm:"many2many:user_rooms;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
