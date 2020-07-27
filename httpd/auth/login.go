package auth

import (
	"encoding/json"
	"net/http"
	// "restaurant-supplier-api/utils/enCors"
)

func Login(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	token, err := GenerateJWT()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(map[string]string{"token": token})

}
