package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;"`
	Name      string    `gorm:"not null;"`
	LastName  string    `gorm:"not null;"`
	Email     string    `gorm:"unique;not null;"`
	Password  string    `gorm:"not null;"`
	Rooms     []Room    `gorm:"many2many:user_rooms;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
