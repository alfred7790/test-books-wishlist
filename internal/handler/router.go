package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-books-wishlist/cmd/config"
	"test-books-wishlist/internal/service"
)

func InitRouter(app *service.Service) *gin.Engine {
	if config.Config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	rootApp := gin.Default()

	rootApp.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Up and running...")
	})

	r := rootApp.Group("v1")
	AddRoutesV1(r, app)
	return rootApp
}
