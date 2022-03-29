package health

import (
	// "encoding/json"
)

type HealthMeasurements struct {
	Data []HealthMeasurement `json:"measurements"`
}

type HealthMeasurement struct {
	Id        			int    `json:"id"`
	CreatedTime 		string `json:"created_time"`
	UpdatedWeight  		string `json:"updated_weight"`
	UpdatedBPSystolic  	string `json:"updated_bp_systolic"`
	UpdatedBPDiastolic  string `json:"updated_bp_diastolic"`
	UpdatedO2 			string `json:"updated_o2"`
	UpdatedBPM  		string `json:"updated_bpm"`
	Comment  			string `json:"comment"`
}