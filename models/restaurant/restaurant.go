package restaurant

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Address  string             `bson:"address,omitempty" json:"address,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone    string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
}
