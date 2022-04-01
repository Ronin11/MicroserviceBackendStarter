package health

import (
	"fmt"
	// "encoding/json"
	"time"

	"nateashby.com/gofun/storage"
)

type HealthMeasurements struct {
	Data []HealthMeasurement `json:"measurements"`
}

type HealthMeasurement struct {
	Id        					int    		`json:"id" db:"id"`
	CreatedTime 				time.Time 	`json:"created_time" db:"created_time"`
	UpdatedWeight  				int 		`json:"updated_weight,omitempty" db:"updated_weight"`
	UpdatedBPSystolic  			int 		`json:"updated_bp_systolic,omitempty" db:"updated_bp_systolic"`
	UpdatedBPDiastolic  		int 		`json:"updated_bp_diastolic,omitempty" db:"updated_bp_diastolic"`
	UpdatedO2 					int 		`json:"updated_o2,omitempty" db:"updated_o2"`
	UpdatedBPM  				int 		`json:"updated_bpm,omitempty" db:"updated_bpm"`
	Comment  					string 		`json:"comment,omitempty" db:"comment"`
}

func GetAllMeasurements() []HealthMeasurement{

	storageHandler := storage.GetInstance()
	fmt.Println("SH2: ", storageHandler)
	storageHandler.Fetch("items")
	var measurements = []HealthMeasurement{}
	
	return measurements
}