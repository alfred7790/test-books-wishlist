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
// @Tags Google Books
// @Produce json
// @Param terms query string true "terms (Author, Title, Publisher)"
// @Param key query string true "Google Books API key"
// @Success 200 {object} entity.BooksDTO
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 404 {object} entity.FailureResponse
// @Router /v1/books/search [GET]
// @Security APIToken
func (service *API) LookForBooks(c *gin.Context) {
	terms := c.DefaultQuery("terms", "")
	apiKey, ok := c.GetQuery("key")
	if !ok {
		common.Failure("We need an APIkey to query google Books API", "Google Books API is empty", http.StatusBadRequest, c)
		return
	}

	resp, err := service.App.GoogleBooksAPI.SearchBooks(terms, apiKey)
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
// @Tags Wish List
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
// @Tags Wish List
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

	response := mapper.WishListsDTOMapper(lists)

	c.JSON(http.StatusOK, response)
}

// AddBookToWishList adding a book into a wishlist
// @Summary adding a book into a wishlist
// @Description adding a book into a wishlist and creating a new book if it doesn't exist
// @Tags Books of Wish List
// @Produce json
// @Param id path int true "WishListID"
// @Param data body entity.Book true "struct to add a book to wishlist"
// @Success 200 {object} entity.SuccessResponse
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 409 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlists/{id}/books [POST]
// @Security APIToken
func (service *API) AddBookToWishList(c *gin.Context) {
	id := c.Param("id")

	wishListId, err := strconv.Atoi(id)
	if err != nil {
		common.Failure("The wishListID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	var input *entity.Book
	err = c.BindJSON(&input)
	if err != nil {
		common.Failure("Wishlist data incorrect", err.Error(), http.StatusBadRequest, c)
		return
	}

	msg, err := service.App.Repo.AddItem(uint(wishListId), input)
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusBadRequest, c)
		return
	}

	c.JSON(http.StatusOK, common.Success())
}

// RemoveItemFromWishList removing a book from wishlist
// @Summary removing a book from wishlist
// @Description removing a book from wishlist if it exists
// @Tags Books of Wish List
// @Produce json
// @Param id path int true "WishListID"
// @Param bookid path string true "BookID"
// @Success 200 {object} entity.SuccessResponse
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlists/{id}/books/{bookid} [DELETE]
// @Security APIToken
func (service *API) RemoveItemFromWishList(c *gin.Context) {
	id := c.Param("id")

	wishListId, err := strconv.Atoi(id)
	if err != nil {
		common.Failure("The wishListID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	bookID := c.Param("bookid")

	_, msg, err := service.App.Repo.GetItem(uint(wishListId), bookID)
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusBadRequest, c)
		return
	}

	msg, err = service.App.Repo.RemoveItem(uint(wishListId), bookID)
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusBadRequest, c)
		return
	}

	c.JSON(http.StatusOK, common.Success())
}

// GetWishList get wishlist if exist
// @Summary returns wishlist of books
// @Description List of books from google saved in datastore
// @Tags Wish List
// @Produce json
// @Param id path int true "WishListID"
// @Success 200 {object} entity.WishListDTO
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlists/{id} [GET]
// @Security APIToken
func (service *API) GetWishList(c *gin.Context) {
	userId, err := common.GetUserIdFromToken(c.Request)
	if err != nil {
		common.Failure("cannot get userId from token", err.Error(), http.StatusInternalServerError, c)
		return
	}

	list, msg, err := service.App.Repo.GetBooksByWishList(userId)
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusInternalServerError, c)
		return
	}

	response := mapper.WishListDTOMapper(list)

	c.JSON(http.StatusOK, response)
}

// RemoveWishList removing a complete wishlist
// @Summary removing a complete wishlist
// @Description all books from a wishlist will be removed
// @Tags Wish List
// @Produce json
// @Param id path int true "WishListID"
// @Success 200 {object} entity.SuccessResponse
// @Failure 400 {object} entity.FailureResponse
// @Failure 401 {object} entity.FailureResponse
// @Failure 403 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/wishlists/{id} [DELETE]
// @Security APIToken
func (service *API) RemoveWishList(c *gin.Context) {
	id := c.Param("id")

	wishListId, err := strconv.Atoi(id)
	if err != nil {
		common.Failure("The wishListID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	_, msg, err := service.App.Repo.WishListExists(uint(wishListId))
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusBadRequest, c)
		return
	}

	msg, err = service.App.Repo.RemoveWishList(uint(wishListId))
	if err != nil {
		common.Failure(msg, err.Error(), http.StatusBadRequest, c)
		return
	}

	c.JSON(http.StatusOK, common.Success())
}
