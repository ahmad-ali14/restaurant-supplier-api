package order

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	From_Restaurant string             `bson:"restaurantId" json:"restaurantId"`
	To_Supplier     string             `bson:"supplierId" json:"supplierId"`
	Status          string             `bson:"status" json:"status"`
	Value           float64            `bson:"value" json:"value"`
	Comments        []*Comment         `bson:"comments" json:"comments"`
}

type Comment struct {
	From    string `bson:"from" json:"from"`
	Message string `bson:"message" json:"message"`
}
