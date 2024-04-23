package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	DeviceId   string               `bson:"device_id,omitempty"`
	Name       string               `bson:"name"`
	RoomID     *primitive.ObjectID  `bson:"room_id"`
	Online     bool                 `bson:"online,omitempty"`
	LastOnline time.Time            `bson:"last_online,omitempty"`
	Configured bool                 `bson:"configured,omitempty"`
	Attributes map[string]Attribute `bson:"attributes"`
	CreatedAt  time.Time            `bson:"created_at,omitempty"`
	UpdatedAt  time.Time            `bson:"updated_at,omitempty"`
}

type Attribute struct {
	Value     string    `bson:"value"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}
