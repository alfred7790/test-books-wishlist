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

// LoginUser login
// @Summary returns a new token
// @Description Used login and get a new token
// @Tags Users
// @Produce json
// @Param data body entity.User true "username and password"
// @Success 200 {object} entity.UserTokenDTO
// @Failure 400 {object} entity.FailureResponse
// @Failure 409 {object} entity.FailureResponse
// @Failure 500 {object} entity.FailureResponse
// @Router /v1/login [POST]
func (service *API) LoginUser(c *gin.Context) {
	var input *entity.User
	err := c.BindJSON(&input)
	if err != nil {
		common.Failure("User data incorrect", err.Error(), http.StatusBadRequest, c)
		return
	}

	input.Password = common.HashString(input.Password)

	user, err := service.App.Repo.UserAuth(input.UserName, input.Password)
	if err != nil {
		if repository.IsRecordNotFound(err) {
			common.Failure("Invalid credentials", err.Error(), http.StatusNotFound, c)
			return
		}
		common.Failure("Error trying to get user data", err.Error(), http.StatusInternalServerError, c)
		return
	}

	token, err := common.GetToken(user.Id, user.UserName)
	if err != nil {
		common.Failure("Error trying to build new token", err.Error(), http.StatusInternalServerError, c)
		return
	}

	data := entity.UserTokenDTO{
		UserID:       user.Id,
		Username:     user.UserName,
		Token:        token,
		TokenSwagger: fmt.Sprintf("Bearer %s", token),
	}

	c.JSON(http.StatusOK, data)
}
