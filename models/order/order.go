package order

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	From_Restaurant string             `bson:"restaurantId,omitempty" json:"restaurantId,omitempty"`
	To_Supplier     string             `bson:"supplierId,omitempty" json:"supplierId,omitempty"`
	Status          string             `bson:"status,omitempty" json:"status,omitempty"`
	Value           float64            `bson:"value,omitempty" json:"value,omitempty"`
	Comments        []*Comment         `bson:"comments,omitempty" json:"comments,omitempty"`
}

type Comment struct {
	From    string `bson:"password,omitempty" json:"password,omitempty"`
	Message string `bson:"password,omitempty" json:"password,omitempty"`
}
