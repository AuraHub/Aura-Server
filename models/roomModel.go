package models

type Room struct {
	ID         uint `gorm:"primaryKey"`
	Name string
	Users []*User `gorm:"many2many:user_rooms;"`
	Devices []Device
}