package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceTrigger struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	DeviceTriggerId string              `bson:"device_trigger_id,omitempty"`
	Name            string              `bson:"name"`
	RoomID          *primitive.ObjectID `bson:"room_id"`
	Online          bool                `bson:"online,omitempty"`
	LastOnline      time.Time           `bson:"last_online,omitempty"`
	Triggers        map[string]Trigger  `bson:"triggers"`
	Configured      bool                `bson:"configured"`
	ConfiguredAt    time.Time           `bson:"configured_at,omitempty"`
	CreatedAt       time.Time           `bson:"created_at,omitempty"`
	UpdatedAt       time.Time           `bson:"updated_at,omitempty"`
}
type Trigger struct {
	Actions  []Action  `bson:"actions"`
	CalledAt time.Time `bson:"called_at,omitempty"`
}

type Action struct {
	DeviceId string `bson:"device_id,omitempty"`
	Action   string `bson:"action,omitempty"`
	Value    string `bson:"value"`
}
