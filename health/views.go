package health

import (
	// "log"
	"fmt"
)

func GetMeasurements() ([]HealthMeasurement, error) {
	storageHandler := storage.GetInstance()
	fmt.Println("SH2: ", storageHandler)

	measurements := GetAllMeasurements()
	

	fmt.Println("WUT2: ", measurements)

	return measurements, nil
}
