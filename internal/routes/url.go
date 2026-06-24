package routes

import (
	"aka-somix/micro-url-shortener/configs"
	"aka-somix/micro-url-shortener/internal/models"
	"aka-somix/micro-url-shortener/internal/repositories"
	"aka-somix/micro-url-shortener/internal/services"
	render "aka-somix/micro-url-shortener/pkg"
	"aka-somix/micro-url-shortener/views/pages"
	"aka-somix/micro-url-shortener/views/pages/home_sections"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type URLsRouter struct {
	urlShortenService *services.UrlShortenService
}

func NewUrlsRouter() (*URLsRouter, error) {
	// urlRepo, err := repositories.NewJsonUrlRepository("tmp/local/database.json")
	urlRepo, err := repositories.NewRedisUrlRepository()

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
				log.Printf("[url] error fetching all URLs: %s", err)
				render.RenderTemplate(c, 503, pages.ServiceUnavailable())
				return
			}
			render.RenderTemplate(c, 200, pages.UrlsAvailable(urls))
		})

		urlGroup.GET("/latest", func(c *gin.Context) {
			n := 5
			if nStr := c.Query("n"); nStr != "" {
				if parsed, err := strconv.Atoi(nStr); err == nil && parsed > 0 {
					n = parsed
				}
			}
			urls, err := router.urlShortenService.GetLatestURLs(n)
			if err != nil {
				log.Printf("[url] error fetching latest: %s", err)
				render.RenderTemplate(c, 500, pages.UrlError("Failed to fetch latest URLs"))
				return
			}
			render.RenderTemplate(c, 200, home_sections.ShortenedUrlRows(urls))
		})

		// POST url/
		// Creates a new url
		urlGroup.POST("/", func(c *gin.Context) {
			var req models.NewURLRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				log.Printf("Error binding JSON: %s", err)
				render.RenderTemplate(c, 200, pages.UrlError("Invalid request: expected JSON with a 'url' field"))
				return
			}

			if req.Url == "" {
				render.RenderTemplate(c, 200, pages.UrlError("URL is required"))
				return
			}

			if strings.Contains(req.Url, configs.BaseURL) {
				render.RenderTemplate(c, 200, pages.UrlError("Cannot shorten a URL from this domain"))
				return
			}

			shortCode, err := router.urlShortenService.ShortenURL(req.Url)
			if err != nil {
				if errors.Is(err, models.ErrStorageFull) {
					render.RenderTemplate(c, 200, pages.UrlError("Storage limit reached — try again later"))
					return
				}
				log.Printf("Error shortening URL: %s", err)
				render.RenderTemplate(c, 200, pages.UrlError("Failed to shorten URL"))
				return
			}

			shortUrl := configs.BaseURL + "/url/" + shortCode
			render.RenderTemplate(c, 201, pages.UrlSuccess(shortUrl))
		})

		urlGroup.GET("/:short", func(c *gin.Context) {
			short := c.Param("short")
			log.Printf("[url] short code: %s", short)

			foundURL, err := router.urlShortenService.GetUrlFromShort(short)
			if err != nil {
				log.Printf("[url] error fetching short code %s: %s", short, err)
				render.RenderTemplate(c, 503, pages.ServiceUnavailable())
				return
			}
			if foundURL == nil {
				log.Printf("[url] short code %s not found", short)
				render.RenderTemplate(c, 404, pages.NotFound())
				return
			}

			log.Printf("[url] original url: %s", foundURL.OriginalURL)
			c.Redirect(302, foundURL.OriginalURL)
		})
	}
}
