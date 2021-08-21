package repository

import (
	"context"
	"fmt"
	"goprueba/Services/auth/dto"

	"go.mongodb.org/mongo-driver/bson"
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
