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
	ID uint64 `json:"id"`
	//	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"bottle_name"`
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
	UserId     uint64
}

type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}

var UserTest = User{
	ID:       1,
	Name:     "username",
	Email:    "mmmmm",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

var users = []User{
	{ID: 1, Email: "juan.ruiz@iwo-os.com", Name: "Juan"},
}