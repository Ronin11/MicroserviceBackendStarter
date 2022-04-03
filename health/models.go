package health

import (
	"fmt"
	"encoding/json"
	"time"
	"errors"

	// "database/sql"
    "database/sql/driver"
)

type HealthMeasurements struct {
	Data []HealthMeasurement `json:"measurements"`
}

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

// type CreateMeasurementResponse struct {
// 	Id	string `json:"id"`
// }

// func (hd HealthData) Value() (driver.Value, error) {
//     return json.Marshal(hd)
// }

// func (hd *HealthData) Scan(value interface{}) error {
//     b, ok := value.([]byte)
//     if !ok {
//         return errors.New("type assertion to []byte failed")
//     }

//     return json.Unmarshal(b, &hd)
// }

func (hm HealthMeasurement) Value() (driver.Value, error) {
	fmt.Println("VALUE: ", hm)
    return json.Marshal(hm)
}

func (hm *HealthMeasurement) Scan(value interface{}) error {
	fmt.Println("SCAN: ", value)
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &hm)
}
