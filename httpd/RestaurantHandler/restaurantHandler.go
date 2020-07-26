package restaurantHandler

import (
	"context"
	"encoding/json"
	// "fmt"
	"go.mongodb.org/mongo-driver/bson"
	// "log"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"restaurant-supplier-api/data/restaurant"
	"restaurant-supplier-api/utils/dbHandler"
	"time"
)

var client = dbHandler.CreateConnection()

var coll = client.Database("go").Collection("restaurants")

func Info(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var routesInfo = make(map[string]map[string]string)

	routesInfo["_welcome "] = map[string]string{"info": "resturant-supplier-api endpoints by Ahmad Ali"}

	routesInfo["route /restaurant "] = map[string]string{
		"Available Methods": "CRUD",
		"this endpoint":     "CRUD on restaurnts",
	}

	json.NewEncoder(res).Encode(routesInfo)
}

func GetRestaurants(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var results []restaurant.Restaurant
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var elem restaurant.Restaurant
		cursor.Decode(&elem)
		results = append(results, elem)
	}
	if err := cursor.Err(); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(results)
}

func CreateRestaurant(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var elem restaurant.Restaurant
	_ = json.NewDecoder(req.Body).Decode(&elem)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := coll.InsertOne(ctx, elem)
	json.NewEncoder(res).Encode(result)
}
