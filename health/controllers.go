package health

import (
	// "log"
	// "reflect"
	// "fmt"

)


// func AddMeasurement(measurement HealthMeasurement) error {
func Add() error {


	healthMeasurement := HealthMeasurement{
		Data: HealthData{
		UpdatedWeight: 101,
		UpdatedBPDiastolic: 102,
		UpdatedBPSystolic: 103,
		UpdatedO2: 104,
		UpdatedBPM: 105,
	}}

	return AddMeasurement(healthMeasurement)
}