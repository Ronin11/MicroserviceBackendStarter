package health

import (
	// "log"
	// "fmt"
)

func GetMeasurements() (*HealthMeasurements, error) {
	storageHandler := GetStorageHandlerInstance()

	return storageHandler.GetAllMeasurements()
}

func GetMeasurement(id string) (*HealthMeasurement, error) {
	storageHandler := GetStorageHandlerInstance()

	return storageHandler.GetMeasurement(id)
}
