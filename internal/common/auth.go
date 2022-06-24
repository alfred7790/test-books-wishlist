package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
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

func TokenValidation(r *http.Request) (string, int, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "Restricted access: you need a token, please login", http.StatusForbidden, errors.New("authorization header is empty")
	}

	if len(strings.Split(header, " ")) <= 0 {
		return "Restricted access: Bearer token is required e.g. Bearer <Token>", http.StatusForbidden, errors.New("authorization header is incomplete")
	}

	tokensSplit := strings.Split(header, " ")
	if len(tokensSplit) != 2 {
		return "Restricted access: Bearer token is required e.g. Bearer <Token>", http.StatusForbidden, errors.New("bearer authorization header is required")
	}

	tokens := tokensSplit[1]
	var userToken entity.UserToken
	token, err := jwt.ParseWithClaims(tokens, &userToken, func(token *jwt.Token) (interface{}, error) { return SecretKey, nil })
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return "Expired Token: please login again", http.StatusUnauthorized, err
		}
		return "Invalid Token: please login again", http.StatusUnauthorized, err
	}

	if userToken.ExpiresAt < time.Now().Unix() {
		return "Invalid Token: please login again", http.StatusUnauthorized, err
	}

	if !token.Valid {
		return "Invalid Token: please login again", http.StatusUnauthorized, err
	}

	return "", http.StatusOK, nil
}
