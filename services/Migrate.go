package services

import (
	"isotopes/models"
)

func Migrate() {
	DB.AutoMigrate(&models.Book{})
}
