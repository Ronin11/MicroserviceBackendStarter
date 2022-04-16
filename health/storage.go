package health

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"

	"nateashby.com/gofun/logging"
	"nateashby.com/gofun/auth"
)

type StorageHandler struct {
	dbUrl string
	db *pgx.Conn
	tableName string
}

var storageHandler *StorageHandler

func initialize(fullUrl string, tableName string) (*StorageHandler) {
	storageHandler = &StorageHandler{dbUrl: fullUrl}
	conn, err := pgx.Connect(context.Background(), fullUrl)
	storageHandler.db = conn
	storageHandler.tableName = tableName
	if err != nil {
		logging.Log("Database connection error: ", err)
	}else{
		logging.Log("Database connection established")
	}
	return storageHandler
}

func (sh *StorageHandler) reconnect() (error){
	conn, err := pgx.Connect(context.Background(), sh.dbUrl)
	storageHandler.db = conn
	if err != nil {
		logging.Log("Database reconnection error: ", err)
		return err
	}
	return nil
} 

func GetStorageHandlerInstance() (*StorageHandler) {
	if storageHandler != nil {
		return storageHandler
	}

	return initialize(os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_HEALTH_TABLE_NAME"))
}

func (sh *StorageHandler) Cleanup() (error){
	sh.db.Close(context.Background())
	return nil
}

func (sh *StorageHandler) GetAllMeasurements(user *auth.User) (*HealthMeasurements, error){
	rows, err := sh.db.Query(context.Background(), fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", sh.tableName), user.Id)
	if err != nil {
		logging.Log("Fetch Err: ", err)
		return nil, err
	}
	defer rows.Close()

	measurements := &HealthMeasurements{}
	for rows.Next() {
		var hm HealthMeasurement
		err := rows.Scan(&hm.Id, &hm.UserId, &hm.CreatedTime, &hm.Data)
		if err != nil {
			logging.Log("SCAN ERR: ", err)
		}
		measurements.Data = append(measurements.Data, hm)
	}

	return measurements, nil
}

func (sh *StorageHandler) GetMeasurement(user *auth.User, id string) (*HealthMeasurement, error){
	var hm HealthMeasurement 
	err := sh.db.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND id=$2", sh.tableName), user.Id, id).Scan(&hm.Id, &hm.UserId, &hm.CreatedTime, &hm.Data)
	if err != nil {
		logging.Log("GET FAILED: ", err)
	}

	return &hm, err
}

func (sh *StorageHandler) CreateMeasurement(user *auth.User, data *HealthData) (*HealthMeasurement, error){
	var hm HealthMeasurement 
	err := sh.db.QueryRow(
		context.Background(),
		fmt.Sprintf("INSERT INTO %s (user_id, data) VALUES($1, $2) RETURNING *", sh.tableName), user.Id, data).Scan(&hm.Id, &hm.UserId, &hm.CreatedTime, &hm.Data)
	if err != nil {
		logging.Log("CREATE FAILED: ", err)
	}

	return &hm, err
}

func (sh *StorageHandler) UpdateMeasurement(user *auth.User, hm *HealthMeasurement) (*HealthMeasurement, error){
	var hm2 HealthMeasurement 
	err := sh.db.QueryRow(
		context.Background(),
		fmt.Sprintf("UPDATE %s SET data=$1 WHERE user_id=$2 AND id=$3 RETURNING *", sh.tableName), hm.Data, user.Id, hm.Id).Scan(&hm2.Id, &hm2.UserId, &hm2.CreatedTime, &hm2.Data)
	if err != nil {
		logging.Log("UPDATE FAILED: ", err)
	}

	return &hm2, err
}

func (sh *StorageHandler) DeleteMeasurement(user *auth.User, id string) (error){

	_, err := sh.db.Exec(
		context.Background(),
		fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2", sh.tableName), user.Id, id)
	if err != nil {
		logging.Log("DELETE FAILED: ", err)
	}

	return err
}


