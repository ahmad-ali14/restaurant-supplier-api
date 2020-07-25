package restaurant

import (
	"restaurant-supplier-api/utils/dbHandler"
)

type Restaurant struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type RestaurantList struct {
	Restaurants []Restaurant
}

func (s *RestaurantList) AddNew(sup Restaurant) string {
	s.Restaurants = append(s.Restaurants, sup)
	return dbHandler.AddItemToDb()

}

func (s *RestaurantList) GetAll() []Restaurant {
	return s.Restaurants
}
