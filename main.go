package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restaurant-supplier-api/config"
)

func main() {
	fmt.Println("working")

	var mongoUri = config.GetMongoUrl()
	fmt.Println(mongoUri)
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
