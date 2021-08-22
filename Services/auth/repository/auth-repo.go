package repository

import (
	"context"
	"fmt"
	"goprueba/Services/auth/dto"
	"goprueba/infrastructure"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

// BottlesAll example
func GetUserByEmail(email string) (dto.User, error) {

	var r dto.User

	col := dto.UserCollection()

	fmt.Printf("email: %v\n", email)

	if err := col.FindOne(ctx, bson.M{"email": email}).Decode(&r); err != nil {
		//log.Fatal(err)
		fmt.Printf("error get user: %v\n", err)
		return r, err
	}

	fmt.Printf("user!!!!: %v\n", r)

	return r, nil
}

func CreateUser(user *dto.User) (*mongo.InsertOneResult, error) {

	col := dto.UserCollection()

	fmt.Printf("user: %v\n", user)
	id := infrastructure.GetSeqId(col)

	fmt.Printf("id: %v\n", id)
	user.ID = id
	result, err := col.InsertOne(ctx, user)

	if err != nil {
		fmt.Println("InsertOne ERROR:", err)
		os.Exit(1) // safely exit script on error
	}

	fmt.Printf("created user!!!!: %v\n", result)

	return result, err
}
