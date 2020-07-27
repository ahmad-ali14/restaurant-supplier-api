package auth

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	// "restaurant-supplier-api/models/restaurant"
	// "restaurant-supplier-api/models/supplier"
	"restaurant-supplier-api/models/user"

	"restaurant-supplier-api/utils/dbHandler"
	"time"
)

var client = dbHandler.CreateConnection()

var restaurantsColl = client.Database("go").Collection("restaurants")
var usersColl = client.Database("go").Collection("users")
var suppliersColl = client.Database("go").Collection("suppliers")

type preUser struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

func Login(res http.ResponseWriter, req *http.Request) {

	// setting header
	res.Header().Set("Content-Type", "application/json")

	// generate token
	token, err := GenerateJWT()

	// if wrong with token generation
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// search the user
	var foundUser user.User
	var userAskingLogin preUser
	var userAccount interface{}

	// decode the req Body
	_ = json.NewDecoder(req.Body).Decode(&userAskingLogin)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err = usersColl.FindOne(ctx, bson.D{
		{"email", userAskingLogin.Email},
	}).Decode(&foundUser)

	log.Println("user Asked Loging:  ", userAskingLogin)
	log.Println("found user:  ", foundUser)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		// res.Write([]byte(`{ "message": "` + err + `" }`))
		log.Println(err)
		return
	}

	//userAskingLogin.Password, _ = HashPassword(userAskingLogin.Password)

	if CheckPasswordHash(userAskingLogin.Password, foundUser.Password) {

		// grap user account
		if foundUser.Role == "restaurant" {
			// var userAccount restaurant.Restaurant
			err = restaurantsColl.FindOne(ctx, bson.D{{"_id", foundUser.UserId}}).Decode(&userAccount)
		}
		if foundUser.Role == "supplier" {
			// var userAccount supplier.Supplier
			err = suppliersColl.FindOne(ctx, bson.D{{"_id", foundUser.UserId}}).Decode(&userAccount)

		}

		json.NewEncoder(res).Encode(map[string]interface{}{"token": token, "role": foundUser.Role, "userId": foundUser.UserId, "userAccount": userAccount})

	} else {
		json.NewEncoder(res).Encode(map[string]string{"Error": "wrong Email or Password"})
	}

}
