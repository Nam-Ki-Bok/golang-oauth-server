package maria

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("MARIA_CONNECTION_INFO")), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
}
