package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server Started")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

	})

	router.Run()
}
