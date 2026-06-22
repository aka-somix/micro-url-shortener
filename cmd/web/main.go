package main

import (
	"aka-somix/micro-url-shortener/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.AddToRouter(router)

	router.Run() // listens on 0.0.0.0:8080
}
