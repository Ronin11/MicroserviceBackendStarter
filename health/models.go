package health

import (
	"encoding/json"
	"time"
	"errors"

    "database/sql/driver"
)

type HealthData struct {
	UpdatedWeight  		int 		`json:"updated_weight,omitempty"`
	UpdatedBPSystolic  	int 		`json:"updated_bp_systolic,omitempty"`
	UpdatedBPDiastolic  int 		`json:"updated_bp_diastolic,omitempty"`
	UpdatedO2 			int 		`json:"updated_o2,omitempty"`
	UpdatedBPM  		int 		`json:"updated_bpm,omitempty"`
	Comment  			string 		`json:"comment,omitempty"`
}

type HealthMeasurement struct {
	Id        	string    	`json:"id"`
	CreatedTime	time.Time 	`json:"created_time"`
	Data		HealthData 	`json:"data"`
}

func (hm *HealthMeasurement) Value() (driver.Value, error) {
    return json.Marshal(hm)
}

func (hm HealthMeasurement) Serialize() ([]byte) {
    str, _ := json.Marshal(hm)
	return str
}

func (hm *HealthMeasurement) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &hm)
}

type HealthMeasurements struct {
	Data []HealthMeasurement `json:"measurements"`
}


func (hms *HealthMeasurements) Value() (driver.Value, error) {
    return json.Marshal(hms)
}

func (hms *HealthMeasurements) Serialize() ([]byte) {
    str, _ := json.Marshal(hms)
	return str
}


