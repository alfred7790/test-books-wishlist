package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-books-wishlist/internal/common"
	"test-books-wishlist/internal/entity"
	"test-books-wishlist/internal/mapper"
	"test-books-wishlist/internal/repository"
)

// LookForBooks get books from google
// @Summary returns a list of books from google books
// @Description List of books from google
// @Tags Books
// @Produce json
// @Param terms query string true "terms (Author, Title, Publisher)"
// @Success 200 {object} entity.BooksDTO
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 404 {object} entity.FailureResponse
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

// CreateWishList create or update a wish list for books
// @Summary create or update a wish list for books
// @Description Used to create a wish list
// @Tags Books
// @Produce json
// @Param userid path string true "userId"
// @Param data body []entity.Book true "struct to create a new wishlist"
// @Success 201 {object} entity.SuccessResponse
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 409 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlist/{userid} [POST]
// @Security APIToken
func (service *API) CreateWishList(c *gin.Context) {
	var input []*entity.Book

	userid := c.Param("userid")
	id, err := strconv.Atoi(userid)
	if err != nil {
		common.Failure("The UserID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	err = c.BindJSON(&input)
	if err != nil {
		common.Failure("Wishlist data incorrect", err.Error(), http.StatusBadRequest, c)
		return
	}

	msg, err := service.App.Repo.CreateWishList(uint(id), input)
	if err != nil {
		if repository.ValidatePSQLError(err, repository.ForeignKeyViolation) {
			common.Failure("Wishlist data incorrect, please review that userID and bookId are correct", err.Error(), http.StatusBadRequest, c)
			return
		}

		common.Failure(msg, err.Error(), http.StatusInternalServerError, c)
		return
	}

	c.JSON(http.StatusOK, common.Success())
}

// GetWishList get wishlist if exists
// @Summary returns wishlist of books
// @Description List of books from google saved in datastore
// @Tags Books
// @Produce json
// @Param userid path string true "userId"
// @Success 200 {object} []entity.ItemWishList
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlist/{userid} [GET]
// @Security APIToken
func (service *API) GetWishList(c *gin.Context) {
	userid := c.Param("userid")
	id, err := strconv.Atoi(userid)
	if err != nil {
		common.Failure("The UserID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	list, msg, err := service.App.Repo.GetWishList(uint(id))
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusInternalServerError, c)
		return
	}

	c.JSON(http.StatusOK, list)
}
