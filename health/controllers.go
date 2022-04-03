package health

import (
	// "log"
	// "reflect"
	// "fmt"

)


// func AddMeasurement(measurement HealthMeasurement) error {
func AddMeasurement(hd HealthData) (HealthMeasurement, error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.CreateMeasurement(hd)
}