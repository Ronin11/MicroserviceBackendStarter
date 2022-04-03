package health

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateHealthRoutes(router *mux.Router) http.Handler {
	log.Println("Creating health routes")
	
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/add", add).Methods("POST")
	
	return router
}

func health(w http.ResponseWriter, r *http.Request) {
	measurements, err := getMeasurements()
	if err != nil {
		log.Println("ERR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(measurements)
	if err != nil {
		log.Println("ERR: ", err)
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
		fmt.Println("ADD ERR: ", err)
	}
	measurement, err := AddMeasurement(entry)
	
	fmt.Println("NEW ITEM ID: ", measurement)
	if err != nil {
		log.Println("ERR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(measurement)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}