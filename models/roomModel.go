package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	Name      string    `gorm:"not null;"`
	Devices   []Device  `gorm:"foreignKey:RoomID"`
	CreatedBy uuid.UUID `gorm:"type:uuid;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
