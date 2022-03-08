package maria

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open(os.Getenv("MARIA_CONNECTION_INFO")), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
	log.Println("Maria connected successfully")
}
