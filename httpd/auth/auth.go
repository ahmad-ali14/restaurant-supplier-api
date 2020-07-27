package auth

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"restaurant-supplier-api/utils/enCors"
	"time"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Hour * 300).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func Login(res http.ResponseWriter, req *http.Request) {

	// res.Header().Set("Content-Type", "application/json")
	// res.Header().Set("Access-Control-Allow-Origin", "*")
	// res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	enCors.EnableCors(res)

	token, err := GenerateJWT()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(map[string]string{"token": token})

}
