package health

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"nateashby.com/gofun/logging"
)

func CreateHealthRoutes(router *mux.Router) http.Handler {
	logging.Log("Creating health routes")
	
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/add", add).Methods("POST")
	
	return router
}

func health(w http.ResponseWriter, r *http.Request) {
	measurements, err := getMeasurements()
	if err != nil {
		logging.Log("ERR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(measurements)
	if err != nil {
		logging.Log("ERR: ", err)
		return
	}

	w.Write(jsonResponse)
}

func add(w http.ResponseWriter, r *http.Request) {
	var entry HealthData
	err := json.NewDecoder(r.Body).Decode(&entry)
	if (HealthData{}) == entry {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("ADD ERR: ", err)
	}
	measurement, err := AddMeasurement(entry)
	
	if err != nil {
		logging.Log("ERR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(measurement)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}