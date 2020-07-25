package restaurantHandler

import (
	"encoding/json"
	"net/http"
)

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
