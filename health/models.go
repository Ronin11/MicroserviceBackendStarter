package health

import (
	// "encoding/json"
)

type Response struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}