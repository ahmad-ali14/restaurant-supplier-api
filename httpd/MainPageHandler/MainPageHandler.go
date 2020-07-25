package mainPageHandler

import (
	"encoding/json"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var routesInfo = make(map[string]map[string]string)

	routesInfo["_welcome "] = map[string]string{"info": "resturant-supplier-api endpoints by Ahmad Ali"}

	routesInfo["route / "] = map[string]string{
		"Available Methods": "GET only",
		"this endpoint":     "will give you more info about the Available routes by this api",
		"more info":         "no authentication yet",
	}

	routesInfo["route /restaurant "] = map[string]string{
		"Available Methods": "CRUD",
		"this endpoint":     "CRUD on restaurnts",
	}

	routesInfo["route /supplier "] = map[string]string{
		"Available Methods": "CRUD",
		"this endpoint":     "CRUD on suppliers",
	}

	routesInfo["route /order "] = map[string]string{
		"Available Methods": "CRUD",
		"this endpoint":     "CRUD on orders",
	}

	json.NewEncoder(w).Encode(routesInfo)
}
