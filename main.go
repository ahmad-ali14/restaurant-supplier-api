package main

import (
	// "context"
	// "fmt"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	// "restaurant-supplier-api/config"
	"restaurant-supplier-api/httpd/mainPageHandler"
	"restaurant-supplier-api/httpd/orderHandler"
	"restaurant-supplier-api/httpd/restaurantHandler"
	"restaurant-supplier-api/httpd/supplierHandler"
	// "restaurant-supplier-api/utils/dbHandler"
	// "time"
)

// client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

func main() {

	// cl := dbHandler.CreateConnection()
	//fmt.Println(mongoURI, &collection)
	// fmt.Printf("client is %s", cl)

	// init router and sub routers
	var router = mux.NewRouter()
	var restaurantRouter = router.PathPrefix("/restaurant").Subrouter()
	var supplierRouter = router.PathPrefix("/supplier").Subrouter()
	var orderRouter = router.PathPrefix("/order").Subrouter()

	// handler funcs
	router.HandleFunc("/", mainPageHandler.Info).Methods("GET")

	// retaurant sub router funcs
	router.Handle("/", restaurantRouter)
	restaurantRouter.HandleFunc("/", restaurantHandler.Info).Methods("GET")
	restaurantRouter.HandleFunc("/all", restaurantHandler.GetRestaurants).Methods("GET")
	restaurantRouter.HandleFunc("/new", restaurantHandler.CreateRestaurant).Methods("POST")


	// supplier sub router funcs
	router.Handle("/", supplierRouter)
	supplierRouter.HandleFunc("/", supplierHandler.Info).Methods("GET")

	// router sub router funcs
	router.Handle("/", orderRouter)
	orderRouter.HandleFunc("/", orderHandler.Info).Methods("GET")

	// run the server
	log.Fatal(http.ListenAndServe(":3000", router))
}
