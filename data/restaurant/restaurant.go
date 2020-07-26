package restaurant

import (
	// "restaurant-supplier-api/utils/dbHandler"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
}

// type RestaurantList struct {
// 	Restaurants []Restaurant
// }

// func (s *RestaurantList) AddNew(sup Restaurant) string {
// 	s.Restaurants = append(s.Restaurants, sup)
// 	return dbHandler.AddItemToDb()

// }

// func (s *RestaurantList) GetAll() []Restaurant {
// 	return s.Restaurants
// }
