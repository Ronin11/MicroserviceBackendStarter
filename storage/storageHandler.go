package storage

import (
	// "reflect"
	"context"
	"fmt"
	"log"
	"os"
	"encoding/json"

	"github.com/jackc/pgx/v4"
)

type StorageHandler struct {
	dbUrl string
	db *pgx.Conn
}

var storageHandler *StorageHandler

func initialize(fullUrl string) (*StorageHandler) {
	storageHandler = &StorageHandler{dbUrl: fullUrl}
	conn, err := pgx.Connect(context.Background(), fullUrl)
	storageHandler.db = conn
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
func GetInstance() (*StorageHandler) {
	if storageHandler != nil {
		return storageHandler
	}

	return initialize(os.Getenv("POSTGRES_URL"))
}

func (sh *StorageHandler) Cleanup() (error){
	sh.db.Close(context.Background())
	return nil
}

func (sh *StorageHandler) Fetch(itemArray interface{}) (error){
	// fmt.Println("FETCH1 : ", reflect.TypeOf(valueType))
	rows, err := sh.db.Query(context.Background(), "SELECT * from items")
	fmt.Println("ROWS: ", rows)
	if err != nil {
		fmt.Println("Fetch Err: ", err)
		return err
	}
	defer rows.Close()
	
	// var rowSlice []Row
	for rows.Next() {
		// var r Row
		// newItem := interface{}
		// err := rows.Scan(&newItem.id, &newItem.created_time, &newItem.data)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// itemArray = append(itemArray, r)
	}

	return nil
	// err := sh.boltDb.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT"))
	// 	b.ForEach(func(k, v []byte) error {
	// 		fmt.Println(string(k), string(v))
	// 		return nil
	// 	})
	// 	return nil
	// })
	// return err
}

func (sh *StorageHandler) Get(id string) (error){
	fmt.Println("GET")
	return nil
	// err := sh.boltDb.View(func(tx *bolt.Tx) error {
	// 	v := tx.Bucket([]byte("DB")).Get([]byte(id))
	// 	fmt.Println("GET: ", v)
	// 	return nil
	// })
	// return err
}



func (sh *StorageHandler) Store(value interface{}) (error) {
	fmt.Println("STORE")
	if value == nil {
		fmt.Println("Store value nil")
		return nil
	}
	if sh.db == nil {
		err := sh.reconnect()
		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Store Marshal Err: ", err)
		return err
	}
	sh.db.Exec(context.Background(), "INSERT INTO items (data) VALUES($1)", data)
	fmt.Println("STORED")
	return nil
	// err := sh.boltDb.Update(func(tx *bolt.Tx) error {
	// 	err := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT")).Put([]byte(id), value)
	// 	if err != nil {
	// 		return fmt.Errorf("could not insert weight: %v", err)
	// 	}
	// 	return nil
	// })
	// fmt.Println("Added Weight")
	// return err
}

// package storage

// import (
// 	"fmt"
// 	// "log"
// 	// "os"
// 	// "strconv"

// 	"github.com/boltdb/bolt"
// )

	
// type StorageHandler struct {
// 	boltDb *bolt.DB
//   //   host string
//   //   port  int
// 	// user string
// 	// password string
// 	// dbname string
// }

// var storageHandler *StorageHandler

// func initialize() (*StorageHandler) {
// 	storageHandler = &StorageHandler{}
// 	db, err := bolt.Open("test.db", 0600, nil)
// 	if err != nil {
// 		fmt.Println("could not open db, %v", err)
// 		return storageHandler
// 	}

// 	err = db.Update(func(tx *bolt.Tx) error {
// 			root, err := tx.CreateBucketIfNotExists([]byte("DB"))
// 			if err != nil {
// 			return fmt.Errorf("could not create root bucket: %v", err)
// 			}
// 			_, err = root.CreateBucketIfNotExists([]byte("WEIGHT"))
// 			if err != nil {
// 			return fmt.Errorf("could not create weight bucket: %v", err)
// 			}
// 			_, err = root.CreateBucketIfNotExists([]byte("ENTRIES"))
// 			if err != nil {
// 			return fmt.Errorf("could not create days bucket: %v", err)
// 			}
// 			return nil
// 	})

// 	if err != nil {
// 			fmt.Println("could not set up buckets, %v", err)
// 			return storageHandler
// 	}
// 	storageHandler.boltDb = db
// 	fmt.Println("DB Setup Done")
// 	return storageHandler
// }

// func Init() (*StorageHandler) {
// 	fmt.Println("Init")
// 	if storageHandler != nil {
// 		return storageHandler
// 	}

// 	fmt.Println("initializing")

// 	return initialize()
// }

// func (sh *StorageHandler) Fetch() (error){
// 	fmt.Println("FETCH")
// 	err := sh.boltDb.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT"))
// 		b.ForEach(func(k, v []byte) error {
// 			fmt.Println(string(k), string(v))
// 			return nil
// 		})
// 		return nil
// 	})
// 	return err
// }

// func (sh *StorageHandler) Get(id string) (error){
// 	fmt.Println("GET")
// 	err := sh.boltDb.View(func(tx *bolt.Tx) error {
// 		v := tx.Bucket([]byte("DB")).Get([]byte(id))
// 		fmt.Println("GET: ", v)
// 		return nil
// 	})
// 	return err
// }



// func (sh *StorageHandler) Set(id string, value []byte) (error){
// 	fmt.Println("STORE")
// 	err := sh.boltDb.Update(func(tx *bolt.Tx) error {
// 		err := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT")).Put([]byte(id), value)
// 		if err != nil {
// 			return fmt.Errorf("could not insert weight: %v", err)
// 		}
// 		return nil
// 	})
// 	fmt.Println("Added Weight")
// 	return err
// }