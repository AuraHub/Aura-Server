package models

type User struct {
	ID         uint `gorm:"primaryKey"`
	Rooms []*Room `gorm:"many2many:user_rooms;"`
	Name string
	LastName string
	Email string	`gorm:"unique"`
	Password string
}