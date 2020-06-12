package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	name string `json:"url" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/api", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to a Go API",
		})
	})

	router.POST("/api/name", func(context *gin.Context) {
		var person Person
		fmt.Println(context)
		context.BindJSON(&person)
		fmt.Println(person)
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %v", person.name),
		})
	})

	router.Run(":5000")
}
