package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
	UserId   primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
}
