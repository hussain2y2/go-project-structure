package controllers

import (
	"isotopes/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Index(c *gin.Context, db *gorm.DB) {
	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
	return
}

func Book(c *gin.Context, db *gorm.DB) {
	id := c.Params.ByName("id")
	var book models.Book

	if err := db.Where("id = ?", id).First(&book).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"data": book})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil})
	}

	return
}
