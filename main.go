package main

import (
	// "context"
	"fmt"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	// "restaurant-supplier-api/config"
	"restaurant-supplier-api/httpd/auth"
	"restaurant-supplier-api/httpd/mainPageHandler"
	"restaurant-supplier-api/httpd/orderHandler"
	"restaurant-supplier-api/httpd/restaurantHandler"
	"restaurant-supplier-api/httpd/supplierHandler"
	// "restaurant-supplier-api/utils/dbHandler"
	jwt "github.com/dgrijalva/jwt-go"
	// "time"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

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
	router.Handle("/", isAuthorized(mainPageHandler.Info)).Methods("GET")
	router.HandleFunc("/login", auth.Login).Methods("POST")

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
	log.Fatal(http.ListenAndServe(":5000", router))
}
