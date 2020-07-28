package supplierHandler

import (
	"context"
	"encoding/json"
	// "fmt"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"log"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"restaurant-supplier-api/httpd/auth"
	"restaurant-supplier-api/models/supplier"
	"restaurant-supplier-api/models/user"
	"restaurant-supplier-api/utils/dbHandler"
	// "restaurant-supplier-api/utils/enCors"

	"time"
)

var client = dbHandler.CreateConnection()

var suppliersColl = client.Database("go").Collection("suppliers")
var usersColl = client.Database("go").Collection("users")

// info
func Info(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var routesInfo = make(map[string]map[string]string)

	routesInfo["_welcome "] = map[string]string{"info": "resturant-supplier-api endpoints by Ahmad Ali"}

	routesInfo["route /supplier "] = map[string]string{
		"Available Methods": "CRUD",
		"this endpoint":     "CRUD on suppliers",
	}

	json.NewEncoder(res).Encode(routesInfo)
}

func GetAllSuppliers(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var results []supplier.Supplier
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := suppliersColl.Find(ctx, bson.M{})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var elem supplier.Supplier
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

func CreateSupplier(res http.ResponseWriter, req *http.Request) {
	// enCors.EnableCors(res)

	var newUser user.User
	var existedUser interface{}
	var elem supplier.RawSupplier
	var supplierToInsert supplier.Supplier

	_ = json.NewDecoder(req.Body).Decode(&elem)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_ = usersColl.FindOne(context.TODO(), bson.D{
		{"email", elem.Email},
	}).Decode(&existedUser)

	if existedUser == nil {
		elem.Password, _ = auth.HashPassword(elem.Password)

		log.Println(elem)

		supplierToInsert.ID = elem.ID
		supplierToInsert.Name = elem.Name
		supplierToInsert.Email = elem.Email
		supplierToInsert.Role = elem.Role
		supplierToInsert.Phone = elem.Phone
		supplierToInsert.Address = elem.Address
		supplierToInsert.Products = elem.Products

		log.Println("Supplier Beffore Inserting:  ", supplierToInsert)

		result, _ := suppliersColl.InsertOne(context.TODO(), supplierToInsert)

		//time.Sleep(3 * time.Second)

		log.Println("Id inserted restaurant", result.InsertedID.(primitive.ObjectID))

		newUser = user.User{Email: elem.Email, Password: elem.Password, Role: elem.Role, UserId: result.InsertedID.(primitive.ObjectID)}
		resultedUser, err := usersColl.InsertOne(ctx, newUser)
		if err != nil {
			log.Println(err.Error())

		}

		log.Println("user Added", resultedUser)

		json.NewEncoder(res).Encode(result)

	} else {
		json.NewEncoder(res).Encode(map[string]string{"Error": "email existed, go to login"})

	}

}
