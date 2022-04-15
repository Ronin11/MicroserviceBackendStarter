package auth

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"nateashby.com/gofun/logging"
)

type AuthRes struct {
	Token	string	`json:"token"`
}

type AuthReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateAuthRoutes(router *mux.Router) http.Handler {
	logging.Log("Creating auth routes")
	
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/createUser", createUser).Methods("POST")
	// router.HandleFunc("/getMeasurement", getMeasurement).Methods("POST")
	// router.HandleFunc("/deleteMeasurement", deleteMeasurement).Methods("POST")
	
	return router
}

func login(w http.ResponseWriter, r *http.Request) {
	logging.Log("Login")
	var authBody AuthReqBody
	err := json.NewDecoder(r.Body).Decode(&authBody)
	if (AuthReqBody{}) == authBody {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("Login ERR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authHandler := GetAuthHandlerInstance()
	token := authHandler.login(authBody.Username, authBody.Password)
	res := &AuthRes{Token: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		logging.Log("ERR: ", err)
		return
	}

	w.Write(jsonResponse)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var authBody AuthReqBody
	err := json.NewDecoder(r.Body).Decode(&authBody)
	if (AuthReqBody{}) == authBody {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		logging.Log("Login ERR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authHandler := GetAuthHandlerInstance()
	token := authHandler.createUser(authBody.Username, authBody.Password)
	res := &AuthRes{Token: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		logging.Log("ERR: ", err)
		return
	}

	w.Write(jsonResponse)
}
