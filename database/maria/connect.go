package maria

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open("mysql", os.Getenv("DATA_CONNECTION_INFO"))
	if err != nil {
		log.Fatal(err.Error())
	}
	db.LogMode(true)

	DB = db
}
