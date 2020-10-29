package controllers

import (
	"isotopes/models"
	"isotopes/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Fetch All the books
 * @param c gin.Context
 */
func AllBooks(c *gin.Context) {
	var books []models.Book
	services.DB.Find(&books)

	response.Data = books
	response.Message = "Data found."
	response.Status = "success"

	c.JSON(http.StatusOK, response)
	return
}

/**
 * Find single book
 * @param c gin.Context
 */
func FindBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book

	if err := services.DB.Where("id = ?", id).First(&book).Error; err != nil {
		response.Data = nil
		response.Status = "error"
		response.Message = "Data not found!"

		c.JSON(http.StatusBadRequest, response)

		return
	}

	response.Data = book
	response.Status = "success"
	response.Message = "Data found"

	c.JSON(http.StatusOK, response)
}
