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
	db := services.Connect()
	services.Migrate(db)
	gin.SetMode(os.Getenv("MODE"))

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		controllers.Index(c, db)
	})

	router.GET("/book/:id", func(c *gin.Context) {
		controllers.Index(c, db)
	})

	router.Run()
}
