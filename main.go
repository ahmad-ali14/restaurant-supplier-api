package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restaurant-supplier-api/config"
	"restaurant-supplier-api/httpd/mainPageHandler"
)

func main() {
	fmt.Println("working")

	var mongoUri = config.GetMongoUrl()
	fmt.Println(mongoUri)

	router := mux.NewRouter()

	router.HandleFunc("/", mainPageHandler.Info).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
