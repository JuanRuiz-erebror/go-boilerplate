package controller

import (
	"fmt"
	"goprueba/Services/sample/model"
	"goprueba/Services/sample/repository"
	"goprueba/infrastructure"
	"goprueba/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ShowBottle godoc
// @Summary Show a bottle
// @Description get string by ID
// @ID get-string-by-int
// @Tags bottles
// @Accept  json
// @Produce  json
// @Param  id path int true "Bottle ID"
// @Success 200 {object} model.Bottle
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /bottles/{id} [get]
func (c *Controller) ShowBottle(ctx *gin.Context) {
	id := ctx.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	bottle, err := repository.BottleOne(bid)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, bottle)
}

// ListBottles godoc
// @Summary List bottles
// @Description get bottles
// @Tags bottles
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Bottle
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /bottles [get]
func (c *Controller) ListBottles(ctx *gin.Context) {
	bottles, err := repository.BottlesAll()
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, bottles)
}

func (c *Controller) ListBras(ctx *gin.Context) {
	infrastructure.LoadEnv()

	mongoUri := os.Getenv("MONGO_URI")

	fmt.Println("mongo uri 111", mongoUri)
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("iwohe").Collection("Bras")
	findOptions := options.Find()

	var results []model.Bra

	// cur, err := db.Find(context.TODO(), bson.M{}, findOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //Finding multiple documents returns a cursor
	// //Iterate through the cursor allows us to decode documents one at a time
	// if err = cur.All(ctx, &results); err != nil {
	// 	log.Fatal(err)
	// }

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	cur, err := db.Find(ctx, bson.M{}, findOptions)
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

	ctx.JSON(http.StatusOK, results)
}
