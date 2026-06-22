package routes

import (
	"github.com/gin-gonic/gin"
)

type URLsRouter struct{}

func NewUrlsRouter() (*URLsRouter, error) {
	return &URLsRouter{}, nil
}

func (router *URLsRouter) AddRoutesTo(group *gin.RouterGroup) {
	urlGroup := group.Group("/url")
	{
		// GET url/
		// Retrieves all urls
		urlGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"available": []string{}})
		})

		// POST url/
		// Creates a new url
		urlGroup.POST("/", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "URL created"})
		})
	}
}
