package routes

import (
	"aka-somix/micro-url-shortener/internal/models"
	render "aka-somix/micro-url-shortener/pkg"
	"aka-somix/micro-url-shortener/views/pages"

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
			render.RenderTemplate(c, 200, pages.UrlsAvailable([]models.URL{}))
		})

		// POST url/
		// Creates a new url
		urlGroup.POST("/", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "URL created"})
		})
	}
}
