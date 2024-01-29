package initializers

import (
	"Aura-Server/models"
)

func SyncDatabase(){

	DB.AutoMigrate(&models.User{})
}