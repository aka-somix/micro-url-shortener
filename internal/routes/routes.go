package routes

import (
	render "aka-somix/micro-url-shortener/pkg"
	"aka-somix/micro-url-shortener/views/pages"

	"github.com/gin-gonic/gin"
)

func AddToRouter(router *gin.Engine) error {
	group := router.Group("/")

	// URLs
	urlsRouter, err := NewUrlsRouter()
	if err != nil {
		return err
	}
	urlsRouter.AddRoutesTo(group)

	// HomePage Route
	router.GET("/", func(c *gin.Context) {
		render.RenderTemplate(c, 200, pages.Home())
	})

	// About Route
	router.GET("/about", func(c *gin.Context) {
		render.RenderTemplate(c, 200, pages.About())
	})

	// 404
	router.NoRoute(func(c *gin.Context) {
		render.RenderTemplate(c, 404, pages.NotFound())
	})

	return nil
}
