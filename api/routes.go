package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"nateashby.com/gofun/logging"
	"nateashby.com/gofun/health"
	"nateashby.com/gofun/auth"
)


func CreateHealthRoutes(router *mux.Router) http.Handler {
	logging.Log("Creating health routes")
	
	router.HandleFunc("/", BuildRouteWithUser(viewMeasurements).Handle).Methods("GET")
	router.HandleFunc("/addMeasurements", BuildRouteWithUser(addMeasurements).Handle).Methods("POST")
	router.HandleFunc("/getMeasurement", BuildRouteWithUser(getMeasurement).Handle).Methods("POST")
	router.HandleFunc("/deleteMeasurement", BuildRouteWithUser(deleteMeasurement).Handle).Methods("POST")
	
	return router
}

func viewMeasurements(w http.ResponseWriter, r *http.Request, user *auth.User) {
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

func addMeasurements(w http.ResponseWriter, r *http.Request, user *auth.User) {
	var entry health.HealthData
	err := json.NewDecoder(r.Body).Decode(&entry)
	if (health.HealthData{}) == entry {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("ADD ERR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	measurement, err := health.AddMeasurement(entry)
	
	if err != nil {
		logging.Log("ERR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(measurement.Serialize())
}

func getMeasurement(w http.ResponseWriter, r *http.Request, user *auth.User) {
	var idObj IdReqObj
	err := json.NewDecoder(r.Body).Decode(&idObj)
	if (IdReqObj{}) == idObj {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("GET ERR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	measurement, err := health.GetMeasurement(idObj.Id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(measurement.Serialize())
}

func updateMeasurement(w http.ResponseWriter, r *http.Request, user *auth.User) {

}

func deleteMeasurement(w http.ResponseWriter, r *http.Request, user *auth.User) {
	var idObj IdReqObj
	err := json.NewDecoder(r.Body).Decode(&idObj)
	if (IdReqObj{}) == idObj {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("DELETE ERR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = health.DeleteMeasurement(idObj.Id)

	if err != nil {
		logging.Log("GET ERR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
}