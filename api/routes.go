package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"nateashby.com/gofun/logging"
	"nateashby.com/gofun/health"
)

func CreateHealthRoutes(router *mux.Router) http.Handler {
	logging.Log("Creating health routes")
	
	router.HandleFunc("/", viewMeasurements).Methods("GET")
	router.HandleFunc("/addMeasurements", addMeasurements).Methods("POST")
	
	return router
}

func viewMeasurements(w http.ResponseWriter, r *http.Request) {
	measurements, err := health.GetMeasurements()
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

func addMeasurements(w http.ResponseWriter, r *http.Request) {
	var entry health.HealthData
	err := json.NewDecoder(r.Body).Decode(&entry)
	if (health.HealthData{}) == entry {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("ADD ERR: ", err)
	}
	measurement, err := health.AddMeasurement(entry)
	
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