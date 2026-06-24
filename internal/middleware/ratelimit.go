package middleware

import (
	render "aka-somix/micro-url-shortener/pkg"
	"aka-somix/micro-url-shortener/views/pages"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit(rps int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(rps), rps)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			render.RenderTemplate(c, http.StatusTooManyRequests, pages.TooManyRequests())
			c.Abort()
			return
		}
		c.Next()
	}
}
