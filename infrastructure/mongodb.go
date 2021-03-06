package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

var ClientMongo *mongo.Client

func MainMongodb() {

	mongoUri := os.Getenv("MONGO_URI")

	fmt.Println("mongo uri 222", mongoUri)
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	ClientMongo = client
}

func GetDB() *mongo.Database {
	db := os.Getenv("MONGO_DB")
	DB := ClientMongo.Database(db)

	return DB
}

func GetSeqId(col *mongo.Collection) int64 {
	res, err := col.EstimatedDocumentCount(ctx)
	if err != nil {
		fmt.Println("seq id ERROR:", err)
		//os.Exit(1) // safely exit script on error
		res = 0
	}
	res++
	return res
}
