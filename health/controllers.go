package health

import (
	// "log"
	"reflect"
	"fmt"

	"nateashby.com/gofun/storage"
)


// func AddMeasurement(measurement HealthMeasurement) error {
func AddMeasurement() error {

	storageHandler := storage.GetInstance()
	fmt.Println("SH: ", storageHandler)

	healthMeasurement := HealthMeasurement{
		UpdatedWeight: 101,
		UpdatedBPDiastolic: 102,
		UpdatedBPSystolic: 103,
		UpdatedO2: 104,
		UpdatedBPM: 105,
	}
	fmt.Println("REF: ", reflect.TypeOf(healthMeasurement))
	return storageHandler.Store(healthMeasurement)
}