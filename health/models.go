package health

import (
	"fmt"
	// "encoding/json"
	"time"

	"nateashby.com/gofun/storage"
	// Grumble
	"upper.io/db.v3/postgresql"
)

type HealthMeasurements struct {
	Data []HealthMeasurement `json:"measurements"`
}

type HealthData struct {
	UpdatedWeight  				int 		`json:"updated_weight,omitempty" db:"updated_weight"`
	UpdatedBPSystolic  			int 		`json:"updated_bp_systolic,omitempty" db:"updated_bp_systolic"`
	UpdatedBPDiastolic  		int 		`json:"updated_bp_diastolic,omitempty" db:"updated_bp_diastolic"`
	UpdatedO2 					int 		`json:"updated_o2,omitempty" db:"updated_o2"`
	UpdatedBPM  				int 		`json:"updated_bpm,omitempty" db:"updated_bpm"`
	Comment  					string 		`json:"comment,omitempty" db:"comment"`
}

type HealthMeasurement struct {
	Id        	int    		`json:"id" db:"id"`
	CreatedTime	time.Time 	`json:"created_time" db:"created_time"`
	Data		HealthData	`db:"data,jsonb"`
	*postgresql.JSONBConverter
}

const healthCollectionConst = "items"

func GetAllMeasurements() []HealthMeasurement{

	storageHandler := storage.GetInstance(healthCollectionConst)
	fmt.Println("sh: ", storageHandler)
	var measurements = []HealthMeasurement{}
	storageHandler.Fetch(measurements)
	fmt.Println("Measurements: ", measurements)
	
	return measurements
}

func AddMeasurement(measurement HealthMeasurement) error {

	storageHandler := storage.GetInstance(healthCollectionConst)
	
	return storageHandler.Store(measurement)
}