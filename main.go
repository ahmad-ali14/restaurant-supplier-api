package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "restaurant-supplier-api/config"
	"restaurant-supplier-api/httpd/mainPageHandler"
	"restaurant-supplier-api/httpd/orderHandler"
	"restaurant-supplier-api/httpd/restaurantHandler"
	"restaurant-supplier-api/httpd/supplierHandler"
)

func main() {

	// mongo conncetion string
	// var mongoURI = config.GetMongoUrl()

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

	// supplier sub router funcs
	router.Handle("/", supplierRouter)
	supplierRouter.HandleFunc("/", supplierHandler.Info).Methods("GET")

	// router sub router funcs
	router.Handle("/", orderRouter)
	orderRouter.HandleFunc("/", orderHandler.Info).Methods("GET")

	// run the server
	log.Fatal(http.ListenAndServe(":3000", router))
}
