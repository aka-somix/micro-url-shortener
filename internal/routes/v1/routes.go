package v1

import (
	"github.com/gin-gonic/gin"
)

func AddRoutesTo(router *gin.Engine) error {
	v1Group := router.Group("/v1")

	// URLs management
	urlRouter, err := NewUrlsRouter()
	if err != nil {
		return err
	}
	urlRouter.AddRoutesTo(v1Group)

	// URL proxy
	proxyRouter, err := NewProxyRouter()
	if err != nil {
		return err
	}
	proxyRouter.AddRoutesTo(v1Group)

	return nil
}
