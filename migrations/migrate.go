package migrations

import (
	"migration/config"
	"migration/models"
)

func Migrate() {
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})
}
