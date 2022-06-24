package entity

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       uint   `json:"-" gorm:"primary_key"`
	UserName string `json:"username" gorm:"not null;unique" example:"Pantufla89" description:"Username"`
	Password string `json:"password" gorm:"not null" example:"Pantufla1234" description:"Password"`
}

type UserDTO struct {
	Id       uint   `json:"id" example:"1" descripcion:"Primary key"`
	UserName string `json:"username" example:"Pantufla89" description:"Username"`
}

type UserToken struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserTokenDTO struct {
	UserID   uint   `json:"userID" example:"1" description:"userID"`
	Username string `json:"username" example:"Pantufla89" description:"Username"`
	Token    string `json:"token" example:"eyJhbGciOiJ..." description:"token for API"`
}
