package repository

import (
	"context"
	"fmt"
	"goprueba/Services/sample/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

// BottlesAll example
func BrasAll() ([]model.Bra, error) {
	findOptions := options.Find()

	var results []model.Bra

	col := model.BraCollection()

	cur, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time
	if err = cur.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	// for cur.Next(ctx) {
	// 	//Create a value into which the single document can be decoded
	// 	var elem model.Bra
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, elem)

	// }

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))

	return results, nil
}
