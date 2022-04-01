package storage

import (
	"fmt"
	"log"
	"os"

	"upper.io/db.v3/postgresql"
	"upper.io/db.v3/lib/sqlbuilder"
)

type DbConfig struct {
	user		string
	password	string
	address		string
	database	string
}

type StorageHandler struct {
	connectionUrl	*postgresql.ConnectionURL
	sess 			sqlbuilder.Database
	collection 		string
	
}

type StorageManager struct {
	handlers map[string]*StorageHandler
}




var storageManager *StorageManager

func initialize(dsn string) (*StorageHandler) {
	storageHandler := &StorageHandler{}
	settings, err := postgresql.ParseURL(dsn)
	if err != nil {
		fmt.Println("DSN Parse ERR: ", err)
	}
	storageHandler.connectionUrl = &settings
	sess, err := postgresql.Open(storageHandler.connectionUrl)
	storageHandler.sess = sess
	if err != nil {
		log.Println("Database connection error: ", err)
	}else{
		log.Println("Database connection established")
	}
	return storageHandler
}

func (sh *StorageHandler) reconnect() (error){
	sess, err := postgresql.Open(sh.connectionUrl)
	sh.sess = sess
	if err != nil {
		log.Println("Database connection error: ", err)
		return err
	}
	return nil
} 

// TODO pass a string for DB Name? Have multiple handlers for each item maybe? or just set it in the struct
func GetInstance(collection string) (*StorageHandler) {
	if storageManager == nil {
		storageManager = &StorageManager{}
		storageManager.handlers = make(map[string]*StorageHandler)
	}
	
	if storageManager.handlers[collection] == nil {
		storageManager.handlers[collection] = initialize(os.Getenv("POSTGRES_URL"))
		storageManager.handlers[collection].collection = collection
	}

	return storageManager.handlers[collection]
}

// func (sh *StorageHandler) Cleanup() (error){
// 	sh.db.Close(context.Background())
// 	return nil
// }

func (sh *StorageHandler) Fetch(items interface{}) (error){
	res := sh.sess.Collection(sh.collection).Find()
	err := res.All(&items)
	return err
}

// func (sh *StorageHandler) Get(id string) (error){
// 	fmt.Println("GET")
// 	return nil
// 	// err := sh.boltDb.View(func(tx *bolt.Tx) error {
// 	// 	v := tx.Bucket([]byte("DB")).Get([]byte(id))
// 	// 	fmt.Println("GET: ", v)
// 	// 	return nil
// 	// })
// 	// return err
// }



func (sh *StorageHandler) Store(value interface{}) (error) {
	fmt.Println("STORE")
	if value == nil {
		fmt.Println("Store value nil")
		return nil
	}
	if sh.sess == nil {
		err := sh.reconnect()
		if err != nil {
			return err
		}
	}
	id, err := sh.sess.Collection(sh.collection).Insert(value)
	if id != nil {
		fmt.Println("STORED")
	}
	return err
}
