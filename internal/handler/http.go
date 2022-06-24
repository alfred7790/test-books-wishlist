package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-books-wishlist/internal/service"
)

type API struct {
	App *service.Service
}

func AddRoutesV1(r *gin.RouterGroup, app *service.Service) {
	/*api := API{
		App: app,
	}*/

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "V1 is running")
	})

	// books := r.Group("books")

}
