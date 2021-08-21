package model

import (
	"goprueba/infrastructure"

	"go.mongodb.org/mongo-driver/mongo"
)

type BraMetadata struct {
	Other string `json:"other" example:"bbbb"`
}

type Bra struct {
	ID       int         `json:"id" example:"1"`
	Name     string      `json:"name" example:"aaaa"`
	Metadata BraMetadata `json:"metadata"`
}

func BraCollection() *mongo.Collection {
	Db := infrastructure.GetDB()
	col := Db.Collection("Bras")
	return col
}
