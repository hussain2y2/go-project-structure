// services/database.go

package services

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	driver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	connection, err := gorm.Open(driver, fmt.Sprint(username+":"+password+"@/"+database+"?charset=utf8&parseTime=True&loc=Local"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = connection
}
