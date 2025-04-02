package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")

	// Retry connection in case Go server is deployed faster than DB server
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connected to DB!")
			return
		}

		fmt.Println("Failed to connect to DB, retrying in 2 seconds...", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Conneciton to DB failed")
	}
}
