package services

import (
	"isotopes/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
}
