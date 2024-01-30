package initializers

import (
	"Aura-Server/models"
)

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Device{},
		&models.Room{},
		&models.Attribute{},
		&models.AttributeValue{},
	)

}
