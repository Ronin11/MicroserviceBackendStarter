package health

import (
	// "log"
	// "fmt"
)

func getMeasurements() (*HealthMeasurements, error) {
	storageHandler := GetStorageHandlerInstance()

	return storageHandler.GetAllMeasurements()
}
