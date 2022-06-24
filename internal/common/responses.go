package common

import (
	"github.com/gin-gonic/gin"
	"test-books-wishlist/internal/entity"
)

const (
	failMessage = "failure"
)

func Failure(message string, details string, code int, ctx *gin.Context) {
	failure := entity.FailureResponse{
		Message: message,
		Status:  failMessage,
		Details: details,
	}
	ctx.JSON(code, failure)
}

// Success is the basic successful HTTP response
func Success() entity.SuccessResponse {
	return entity.SuccessResponse{
		Success: true,
		Message: "ok",
	}
}
