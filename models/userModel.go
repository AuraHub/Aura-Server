package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	LastName string
	Email string	`gorm:"unique"`
	Password string
}