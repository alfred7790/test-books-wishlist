package handler

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"test-books-wishlist/cmd/config"
	"test-books-wishlist/internal/entity"
	"time"
)

var SecretKey []byte

func init() {
	SecretKey = []byte(config.Config.SecretKey)
}

func GetToken(userID uint, username string) (string, error) {
	data := jwt.New(jwt.SigningMethodHS256)
	claims := entity.UserToken{}
	claims.UserID = userID
	claims.Username = username
	claims.Id = strconv.Itoa(int(userID))
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(12)).Unix()
	claims.IssuedAt = time.Now().Unix()

	data.Claims = claims

	token, err := data.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
