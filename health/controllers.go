package health


func AddMeasurement(hd HealthData) (HealthMeasurement, error) {
	storageHandler := GetStorageHandlerInstance()
	return storageHandler.CreateMeasurement(hd)
}