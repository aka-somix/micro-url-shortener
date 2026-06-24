package main

import (
	"aka-somix/micro-url-shortener/configs"
	"aka-somix/micro-url-shortener/internal/middleware"
	"aka-somix/micro-url-shortener/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.RateLimit(configs.RateLimitRPS))

	error := routes.AddToRouter(router)
	if error != nil {
		panic(error)
	}

	router.Run() // listens on 0.0.0.0:8080
}
