package health

import (
	// "reflect"
	"context"
	"fmt"
	"log"
	"os"
	// "encoding/json"

	"github.com/jackc/pgx/v4"
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
		log.Println("Database connection error: ", err)
	}else{
		log.Println("Database connection established")
	}
	return storageHandler
}

func (sh *StorageHandler) reconnect() (error){
	conn, err := pgx.Connect(context.Background(), sh.dbUrl)
	storageHandler.db = conn
	if err != nil {
		log.Println("Database reconnection error: ", err)
		return err
	}
	return nil
} 

// TODO pass a string for DB Name? Have multiple handlers for each item maybe? or just set it in the struct
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

func (sh *StorageHandler) GetAllMeasurements() (*HealthMeasurements, error){
	rows, err := sh.db.Query(context.Background(), fmt.Sprintf("SELECT * from %s", sh.tableName))
	if err != nil {
		fmt.Println("Fetch Err: ", err)
		return nil, err
	}
	defer rows.Close()

	measurements := &HealthMeasurements{}
	for rows.Next() {
		var hm HealthMeasurement
		err := rows.Scan(&hm.Id, &hm.CreatedTime, &hm.Data)
		if err != nil {
			fmt.Println("SCAN ERR: ", err)
			// log.Fatal(err)
		}
		measurements.Data = append(measurements.Data, hm)
	}

	return measurements, nil
}

func (sh *StorageHandler) CreateMeasurement(data HealthData) (HealthMeasurement, error){
	var hm HealthMeasurement 
	err := sh.db.QueryRow(context.Background(), fmt.Sprintf("INSERT INTO %s (data) VALUES($1) RETURNING *", sh.tableName), data).Scan(&hm.Id, &hm.CreatedTime, &hm.Data)
	if err != nil {
		fmt.Println("CREATE FAILED: ", err)
	}

	return hm, err
}

func (sh *StorageHandler) Get(id string) (error){
	fmt.Println("GET")
	return nil
}



// func (sh *StorageHandler) Store(value interface{}) (error) {
// 	fmt.Println("STORE")
// 	if value == nil {
// 		fmt.Println("Store value nil")
// 		return nil
// 	}
// 	if sh.db == nil {
// 		err := sh.reconnect()
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	data, err := json.Marshal(value)
// 	if err != nil {
// 		fmt.Println("Store Marshal Err: ", err)
// 		return err
// 	}
	
// 	fmt.Println("STORED")
// 	return nil
// }
