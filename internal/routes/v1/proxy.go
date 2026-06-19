package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProxyRouter struct{}

func NewProxyRouter() (*ProxyRouter, error) {
	return &ProxyRouter{}, nil
}

func (router *ProxyRouter) AddRoutesTo(group *gin.RouterGroup) {
	urlGroup := group.Group("/proxy")
	{
		// GET proxy/:short
		// Redirects to the original URL based on the provided short URL ID
		urlGroup.GET("/:short", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "https://example.com")
		})
	}
}
