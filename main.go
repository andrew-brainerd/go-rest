package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to a Go API",
		})
	})

	router.POST("/api/name", func(context *gin.Context) {
		var person Person
		context.BindJSON(&person)
		fmt.Println(person)
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %v", person.Name),
		})
	})

	router.Run("0.0.0.0:5000")
}
