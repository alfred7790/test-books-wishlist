package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "test-books-wishlist/docs"
	"test-books-wishlist/internal/middleware"
	"test-books-wishlist/internal/service"
)

type API struct {
	App *service.Service
}

func AddRoutesV1(r *gin.RouterGroup, app *service.Service) {
	api := API{
		App: app,
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "V1 is running")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := r.Group("users")
	users.POST("", api.CreateUser)

	auth := r.Group("login")
	auth.POST("", api.LoginUser)

	books := r.Group("books").Use(middleware.IsAuth)
	books.GET("", api.LookForBooks)

	wishlist := r.Group("wishlist").Use(middleware.IsAuth)
	wishlist.POST(":userid", api.CreateWishList)
	wishlist.GET(":userid", api.GetWishList)
}
