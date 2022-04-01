package health

import (
	// "log"
	"fmt"
)

func GetMeasurements() ([]HealthMeasurement, error) {

	measurements := GetAllMeasurements()
	

	fmt.Println("WUT2: ", measurements)

	return measurements, nil
}
