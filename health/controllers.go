package health


func AddMeasurement(hd HealthData) (*HealthMeasurement, error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.CreateMeasurement(hd)
}

func DeleteMeasurement(id string) (error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.DeleteMeasurement(id)
}