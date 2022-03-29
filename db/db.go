package db

import (
	// "database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

	
type dbConfig struct {
    host string
    port  int
	user string
	password string
	dbname string
}

var db *sqlx.DB

func Initialize(config dbConfig) (*sqlx.DB, error) {
	// db := sql.Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.host, config.port, config.user, config.password, config.dbname)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return db, err
	}

	// db.Conn = conn
	err = db.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

func GetDB() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	config := dbConfig{
		host: os.Getenv("DB_URL"),
		port: port,
		user: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname: os.Getenv("DB_NAME"),
	}
	return Initialize(config)
}