package initializers

import "github.com/hanno-meister/OAuth2Server_challenge/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
