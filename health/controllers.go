package health

import (
	"nateashby.com/gofun/auth"
)


func AddMeasurement(user *auth.User, hd *HealthData) (*HealthMeasurement, error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.CreateMeasurement(user, hd)
}

func UpdateMeasurement(user *auth.User, hm *HealthMeasurement) (*HealthMeasurement, error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.UpdateMeasurement(user, hm)
}

func DeleteMeasurement(user *auth.User, id string) (error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.DeleteMeasurement(user, id)
}