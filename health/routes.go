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
	
	return router
}

func health(w http.ResponseWriter, r *http.Request) {
	var response Response
	persons := GetPersons()

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}