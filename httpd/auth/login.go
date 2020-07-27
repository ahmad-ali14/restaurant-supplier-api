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
	token, tokenErr := GenerateJWT()

	if tokenErr != nil {
		log.Println(tokenErr)
	}

	// if wrong with token generation
	if tokenErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + tokenErr.Error() + `" }`))
		return
	}

	// search the user
	var foundUser user.User
	var userAskingLogin preUser
	var userAccount interface{}

	// decode the req Body
	_ = json.NewDecoder(req.Body).Decode(&userAskingLogin)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	findUserErr := usersColl.FindOne(ctx, bson.D{
		{"email", userAskingLogin.Email},
	}).Decode(&foundUser)

	log.Println("user Asked Loging:  ", userAskingLogin)
	log.Println("found user:  ", foundUser)

	if findUserErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + findUserErr.Error() + `" }`))
		log.Println(findUserErr)
		return
	}

	//userAskingLogin.Password, _ = HashPassword(userAskingLogin.Password)

	if CheckPasswordHash(userAskingLogin.Password, foundUser.Password) {
		var accountErr interface{}
		// grap user account
		if foundUser.Role == "restaurant" {
			// var userAccount restaurant.Restaurant
			accountErr = restaurantsColl.FindOne(ctx, bson.D{{"_id", foundUser.UserId}}).Decode(&userAccount)
		}
		if foundUser.Role == "supplier" {
			// var userAccount supplier.Supplier
			accountErr = suppliersColl.FindOne(ctx, bson.D{{"_id", foundUser.UserId}}).Decode(&userAccount)

		}

		if accountErr != nil {
			log.Println(accountErr)
		}
		if token == "" {
			json.NewEncoder(res).Encode(map[string]string{"Error": "Token is Worng"})
			return
		}

		var results = map[string]interface{}{"token": token, "role": foundUser.Role, "userId": foundUser.UserId, "userAccount": userAccount}

		log.Println(results)

		json.NewEncoder(res).Encode(results)

	} else {
		json.NewEncoder(res).Encode(map[string]string{"Error": "wrong Email or Password"})
	}

}
