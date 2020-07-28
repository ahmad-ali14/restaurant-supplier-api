package supplier

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Supplier struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`

	// Password string `bson:"password,omitempty" json:"password,omitempty"`

	Name string `bson:"name,omitempty" json:"name,omitempty"`

	Address string `bson:"address,omitempty" json:"address,omitempty"`

	Phone string `bson:"phone,omitempty" json:"phone,omitempty"`

	Role string `bson:"role,omitempty" json:"role,omitempty"`

	Products []*Product `bson:"products,omitempty" json:"products,omitempty"`
}

type RawSupplier struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`

	Password string `bson:"password,omitempty" json:"password,omitempty"`

	Name string `bson:"name,omitempty" json:"name,omitempty"`

	Address string `bson:"address,omitempty" json:"address,omitempty"`

	Phone string `bson:"phone,omitempty" json:"phone,omitempty"`

	Role string `bson:"role,omitempty" json:"role,omitempty"`

	Products []*Product `bson:"products,omitempty" json:"products,omitempty"`
}

type Product struct {
	Name string `bson:"productName,omitempty" json:"productName,omitempty"`

	Price string `bson:"productPrice,omitempty" json:"productPrice,omitempty"`
}
