package restaurant

type Restaurant struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type RestaurantList struct {
	Restaurants []Restaurant
}
