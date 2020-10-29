package main

import (
	"isotopes/controllers"
	"isotopes/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	services.Connect()
	services.Migrate()
	gin.SetMode(os.Getenv("MODE"))

	router := gin.Default()

	router.GET("/books", func(c *gin.Context) {
		controllers.AllBooks(c)
	})

	router.GET("/book/:id", func(c *gin.Context) {
		controllers.FindBook(c)
	})

	router.Run()
}
