package routes

import (
	"aka-somix/micro-url-shortener/internal/repositories"
	"aka-somix/micro-url-shortener/internal/services"
	render "aka-somix/micro-url-shortener/pkg"
	"aka-somix/micro-url-shortener/views/pages"
	"log"

	"github.com/gin-gonic/gin"
)

type URLsRouter struct {
	urlShortenService *services.UrlShortenService
}

func NewUrlsRouter() (*URLsRouter, error) {
	urlRepo, err := repositories.NewJsonUrlRepository("tmp/local/database.json")
	if err != nil {
		return nil, err
	}

	urlShortenService := services.NewUrlShortenService(urlRepo)

	return &URLsRouter{urlShortenService: urlShortenService}, nil
}

func (router *URLsRouter) AddRoutesTo(group *gin.RouterGroup) {
	urlGroup := group.Group("/url")
	{
		// GET url/
		// Retrieves all urls
		urlGroup.GET("/", func(c *gin.Context) {
			urls, err := router.urlShortenService.GetAllURLs()
			if err != nil {
				// TODO: Handle error
				return
			}
			render.RenderTemplate(c, 200, pages.UrlsAvailable(urls))
		})

		// POST url/
		// Creates a new url
		urlGroup.POST("/", func(c *gin.Context) {
			router.urlShortenService.ShortenURL("https://example.com")
		})

		urlGroup.GET("/:short", func(c *gin.Context) {
			short := c.Param("short")
			log.Printf("[url] short code: %s", short)

			foundURL, err := router.urlShortenService.GetUrlFromShort(short)

			log.Printf("[url] original url: %s", foundURL.OriginalURL)

			if err != nil {
				// TODO: Handle error
				render.RenderTemplate(c, 404, pages.NotFound())
			}

			c.Redirect(302, foundURL.OriginalURL)
		})
	}
}
