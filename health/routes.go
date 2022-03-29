package health

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateHealthRoutes(router *mux.Router) http.Handler {
	log.Println("Creating health routes")
	
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/add", add).Methods("GET")
	
	return router
}

func health(w http.ResponseWriter, r *http.Request) {
	measurements, err := GetMeasurements()
	if err != nil {
		log.Println("ERR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(measurements)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func add(w http.ResponseWriter, r *http.Request) {
	err := AddMeasurement()
	if err != nil {
		log.Println("ERR: ", err)
	}

	w.WriteHeader(http.StatusOK)
}