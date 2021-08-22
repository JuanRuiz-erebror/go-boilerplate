package dto

import (
	"goprueba/infrastructure"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type User struct {
	ID int64 `json:"id"`
	//	ID       int    `json:"id" example:"1"`
	Name     string `json:"name"`
	Lastname string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func UserCollection() *mongo.Collection {
	Db := infrastructure.GetDB()
	col := Db.Collection("User")
	return col
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type AccessDetails struct {
	AccessUuid string
	UserId     int64
}
