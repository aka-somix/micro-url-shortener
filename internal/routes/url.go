package routes

import (
	"aka-somix/micro-url-shortener/internal/models"
	render "aka-somix/micro-url-shortener/pkg"
	"aka-somix/micro-url-shortener/views/pages"
	"log"

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

		urlGroup.GET("/:short", func(c *gin.Context) {
			short := c.Param("short")
			log.Printf("[url] short code: %s", short)

			if short == "ex1a2b" {
				c.Redirect(302, "https://example.com")
			} else {
				render.RenderTemplate(c, 404, pages.NotFound())
			}
		})
	}
}
