package health

import (
	// "encoding/json"
	"time"
)

type HealthMeasurements struct {
	Data []HealthMeasurement `json:"measurements"`
}

type HealthMeasurement struct {
	Id        					int    `json:"id"`
	CreatedTime 				time.Time `json:"created_time"`
	UpdatedWeight  			int `json:"updated_weight,omitempty"`
	UpdatedBPSystolic  	int `json:"updated_bp_systolic,omitempty"`
	UpdatedBPDiastolic  int `json:"updated_bp_diastolic,omitempty"`
	UpdatedO2 					int `json:"updated_o2,omitempty"`
	UpdatedBPM  				int `json:"updated_bpm,omitempty"`
	Comment  						string `json:"comment,omitempty"`
}

func GetAllMeasurements() []HealthMeasurement{
	var measurements = []HealthMeasurement{}
	
	return measurements
}