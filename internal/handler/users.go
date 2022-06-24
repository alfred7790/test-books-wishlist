package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test-books-wishlist/internal/common"
	"test-books-wishlist/internal/entity"
	"test-books-wishlist/internal/mapper"
	"test-books-wishlist/internal/repository"
)

// CreateUser create a new user
// @Summary returns details about a new user was created
// @Description Used to create a new user
// @Tags Users
// @Produce json
// @Param data body entity.User true "struct to create a new user"
// @Success 201 {object} entity.UserDTO
// @Failure 400 {object} entity.FailureResponse
// @Failure 409 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/users [POST]
func (service *API) CreateUser(c *gin.Context) {
	var input *entity.User
	err := c.BindJSON(&input)
	if err != nil {
		common.Failure("User data incorrect", err.Error(), http.StatusBadRequest, c)
		return
	}

	input.Password = common.HashString(input.Password)

	err = service.App.Repo.CreateUser(input)
	if err != nil {
		if repository.ValidatePSQLError(err, repository.UniqueViolation) {
			msg := fmt.Sprintf("The username %s already exists", input.UserName)
			common.Failure(msg, err.Error(), http.StatusConflict, c)
			return
		}
		common.Failure("Error trying to create a new user", err.Error(), http.StatusInternalServerError, c)
		return
	}

	created := mapper.UserDTOMapper(input)

	c.JSON(http.StatusCreated, created)
}
