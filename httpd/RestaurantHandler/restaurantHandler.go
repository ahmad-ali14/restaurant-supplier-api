package restaurantHandler

import (
	"context"
	"encoding/json"
	// "fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"restaurant-supplier-api/httpd/auth"
	"restaurant-supplier-api/models/restaurant"
	"restaurant-supplier-api/models/user"
	"restaurant-supplier-api/utils/dbHandler"
	"restaurant-supplier-api/utils/enCors"

	"time"
)

var client = dbHandler.CreateConnection()

var restaurantsColl = client.Database("go").Collection("restaurants")
var usersColl = client.Database("go").Collection("users")

func Info(res http.ResponseWriter, req *http.Request) {

	enCors.EnableCors(res)

	var routesInfo = make(map[string]map[string]string)

	routesInfo["_welcome "] = map[string]string{"info": "resturant-supplier-api endpoints by Ahmad Ali"}

	routesInfo["route /restaurant "] = map[string]string{
		"Available Methods": "CRUD",
		"this endpoint":     "CRUD on restaurnts",
	}

	json.NewEncoder(res).Encode(routesInfo)
}

func GetAllRestaurants(res http.ResponseWriter, req *http.Request) {
	// enCors.EnableCors(res)

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,  Accept-Language, YourOwnHeader")

	if (*req).Method == "OPTIONS" {
		return
	}

	var results []restaurant.Restaurant
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := restaurantsColl.Find(ctx, bson.M{})

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
	enCors.EnableCors(res)
	var elem restaurant.Restaurant
	var newUser user.User
	_ = json.NewDecoder(req.Body).Decode(&elem)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	elem.Password, _ = auth.HashPassword(elem.Password)

	log.Println(elem)

	result, _ := restaurantsColl.InsertOne(ctx, elem)

	log.Println(result.InsertedID.(primitive.ObjectID))

	newUser = user.User{Email: elem.Email, Password: elem.Password, Role: elem.Role, UserId: result.InsertedID.(primitive.ObjectID)}
	resultedUseer, err := usersColl.InsertOne(ctx, newUser)
	if err != nil {
		log.Println(err.Error())

	}

	log.Println("user Added", resultedUseer)

	json.NewEncoder(res).Encode(result)
}
