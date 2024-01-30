package initializers

import (
	"Aura-Server/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Device{})
	DB.AutoMigrate(&models.Room{})
	DB.AutoMigrate(&models.Attribute{})
	DB.AutoMigrate(&models.AttributeValue{})
}
