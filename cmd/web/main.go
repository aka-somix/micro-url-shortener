package main

import (
	v1 "aka-somix/micro-url-shortener/internal/routes/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1.AddRoutesTo(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthy",
		})
	})
	router.Run() // listens on 0.0.0.0:8080
}
