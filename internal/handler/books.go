package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Param data body entity.WishList true "struct to create a new wishlist"
// @Success 201 {object} entity.SuccessResponse
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 409 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlists [POST]
// @Security APIToken
func (service *API) CreateWishList(c *gin.Context) {
	userId, err := common.GetUserIdFromToken(c.Request)
	if err != nil {
		common.Failure("cannot get userId from token", err.Error(), http.StatusInternalServerError, c)
		return
	}

	var input *entity.WishList
	err = c.BindJSON(&input)
	if err != nil {
		common.Failure("Wishlist data incorrect", err.Error(), http.StatusBadRequest, c)
		return
	}

	input.UserID = userId

	msg, err := service.App.Repo.CreateWishList(input)
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

// GetWishLists get wishlists if exist
// @Summary returns wishlist of books
// @Description List of books from google saved in datastore
// @Tags Books
// @Produce json
// @Success 200 {object} []entity.WishListDTO
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlists [GET]
// @Security APIToken
func (service *API) GetWishLists(c *gin.Context) {
	userId, err := common.GetUserIdFromToken(c.Request)
	if err != nil {
		common.Failure("cannot get userId from token", err.Error(), http.StatusInternalServerError, c)
		return
	}

	lists, msg, err := service.App.Repo.GetWishLists(userId)
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusInternalServerError, c)
		return
	}

	response := mapper.WishListMapper(lists)

	c.JSON(http.StatusOK, response)
}
