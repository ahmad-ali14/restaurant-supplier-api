package orderHandler

import (
	"context"
	"encoding/json"
	// "fmt"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "log"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	// "restaurant-supplier-api/httpd/auth"
	"restaurant-supplier-api/models/order"
	// "restaurant-supplier-api/models/user"
	"restaurant-supplier-api/utils/dbHandler"
	// "restaurant-supplier-api/utils/enCors"
	"github.com/gorilla/mux"
	"log"
	"time"
)

var client = dbHandler.CreateConnection()

var ordersColl = client.Database("go").Collection("orders")

func Info(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var routesInfo = make(map[string]map[string]string)

	routesInfo["_welcome "] = map[string]string{"info": "resturant-supplier-api endpoints by Ahmad Ali"}

	routesInfo["route /order "] = map[string]string{
		"Available Methods": "CRUD",
		"this endpoint":     "CRUD on orders",
	}

	json.NewEncoder(res).Encode(routesInfo)
}

func GetAllOrders(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var results []order.Order
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := ordersColl.Find(ctx, bson.M{})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var elem order.Order
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

func CreateOrder(res http.ResponseWriter, req *http.Request) {

	var elem order.Order

	_ = json.NewDecoder(req.Body).Decode(&elem)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, _ := ordersColl.InsertOne(ctx, elem)

	log.Println("order Added")

	json.NewEncoder(res).Encode(result)
}

func GetMyOrders(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	id := params["id"]

	log.Println("id:   ", id)

	docID, err := primitive.ObjectIDFromHex(id)

	log.Println("docID:   ", docID)

	var results []order.Order

	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	ctx := context.TODO()
	cursor, err := ordersColl.Find(ctx,
		bson.M{"$or": []bson.M{bson.M{"restaurantId": id}, bson.M{"supplierId": id}}},
		// bson.M{"restaurantId": id},
	)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	log.Println("cursor:   ", cursor)

	for cursor.Next(ctx) {
		var elem order.Order
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
