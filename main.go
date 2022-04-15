package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"nateashby.com/gofun/api"
	"nateashby.com/gofun/auth"
)

func main() {
	log.Println("Starting server")
	router := mux.NewRouter()
	router.HandleFunc("/heartbeat", Heartbeat).Methods("GET")
    
	healthRouter := router.PathPrefix("/health").Subrouter()
	api.CreateHealthRoutes(healthRouter)
	authRouter := router.PathPrefix("/auth").Subrouter()
	auth.CreateAuthRoutes(authRouter)
	
    http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}