package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-books-wishlist/internal/common"
	"test-books-wishlist/internal/mapper"
)

// LookForBooks get books from google
// @Summary returns a list of books from google books
// @Description List of books from google
// @Tags Books
// @Produce json
// @Param terms query string true "terms (Author, Title, Publisher)"
// @Success 200 {object} entity.BooksDTO
// @Failure 400 {object} entity.FailureResponse
// @Failure 404 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/books [GET]
// @Security APIToken
func (service *API) LookForBooks(c *gin.Context) {
	terms := c.DefaultQuery("terms", "")

	resp, err := service.App.GoogleBooksAPI.SearchBooks(terms)
	if err != nil {
		common.Failure("Cannot request google books at the moment", err.Error(), http.StatusBadRequest, c)
		return
	}

	list := mapper.GoogleBookToBooksDTO(resp)

	c.JSON(http.StatusOK, list)
}
