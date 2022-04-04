package health

import (
	// "log"
	// "fmt"
)

func GetMeasurements() (*HealthMeasurements, error) {
	storageHandler := GetStorageHandlerInstance()

	return storageHandler.GetAllMeasurements()
}
