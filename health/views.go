package health

import (
	"nateashby.com/gofun/auth"
)

func GetMeasurements(user *auth.User) (*HealthMeasurements, error) {
	storageHandler := GetStorageHandlerInstance()

	return storageHandler.GetAllMeasurements(user)
}

func GetMeasurement(user *auth.User, id string) (*HealthMeasurement, error) {
	storageHandler := GetStorageHandlerInstance()

	return storageHandler.GetMeasurement(user, id)
}
