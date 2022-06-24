package middleware

import (
	"github.com/gin-gonic/gin"
	"test-books-wishlist/internal/common"
)

func IsAuth(ctx *gin.Context) {
	msg, code, err := common.TokenValidation(ctx.Request)
	if err != nil {
		common.Failure(msg, err.Error(), code, ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}
