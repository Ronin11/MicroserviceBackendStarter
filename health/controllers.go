package health

import (
	"log"
	"fmt"

	"nateashby.com/golang/db"
	// "github.com/jmoiron/sqlx"
)

func GetMeasurements() (HealthMeasurements, error) {
	db, err := db.GetDB()
	if err != nil {
		log.Println("ERR: ", err)
	}

	measurements := HealthMeasurements{}

	// rows, err := db.Query("SELECT * FROM items ORDER BY ID DESC")
	db.Select(&measurements, "SELECT * FROM items ORDER BY ID DESC")

	fmt.Println("MEASUREMENTS: ", measurements)
	// if err != nil {
	// 	return measurements, err
	// }

	// for rows.Next() {
	// 	var measurement HealthMeasurement
	// 	err := rows.Scan(&measurement.Id, &measurement.Name, &measurement.Description, &measurement.CreatedAt)
	// 	if err != nil {
	// 		return measurements, err
	// 	}
	// 	measurements.Data = append(measurements.Data, measurement)
	// }
	return measurements, nil
}

// func AddMeasurement(measurement HealthMeasurement) error {
func AddMeasurement() error {
	db, dbErr := db.GetDB()
	if dbErr != nil {
		log.Println("ERR: ", dbErr)
	}

	var id int
	var createdTime string
	query := `INSERT INTO
	items (updated_weight, updated_bp_systolic, updated_bp_diastolic, updated_o2, updated_bpm, comment)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_time`
	err := db.QueryRow(query, 100, 101, 102, 103, 104, "COMMENT").Scan(&id, &createdTime)
	if err != nil {
		return err
	}

	// item.ID = id
	// item.CreatedAt = createdAt
	return nil
}