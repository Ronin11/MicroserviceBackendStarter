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
	router.HandleFunc("/getMeasurement", getMeasurement).Methods("POST")
	router.HandleFunc("/deleteMeasurement", deleteMeasurement).Methods("POST")
	
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
	w.Write(measurement.Serialize())
}

func getMeasurement(w http.ResponseWriter, r *http.Request) {
	var idObj IdReqObj
	err := json.NewDecoder(r.Body).Decode(&idObj)
	if (IdReqObj{}) == idObj {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("GET ERR: ", err)
	}
	measurement, err := health.GetMeasurement(idObj.Id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(measurement.Serialize())
}

func updateMeasurement(w http.ResponseWriter, r *http.Request) {

}

func deleteMeasurement(w http.ResponseWriter, r *http.Request) {
	var idObj IdReqObj
	err := json.NewDecoder(r.Body).Decode(&idObj)
	if (IdReqObj{}) == idObj {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("DELETE ERR: ", err)
	}
	err = health.DeleteMeasurement(idObj.Id)

	if err != nil {
		logging.Log("GET ERR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}