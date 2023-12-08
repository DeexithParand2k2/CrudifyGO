package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	r "reflect"

	_ "github.com/go-sql-driver/mysql"
	godot "github.com/joho/godotenv"
)

var db *sql.DB // Declare a global variable for the database connection pool

// from env
type DbConfig struct {
	Username     string
	Password     string
	Network      string
	DatabasePort string // Default PORT for MySQL
	Host         string // Localhost IP Addr
}

// load local .env : error handling done
func LoadEnv() error {

	err := godot.Load(".env")
	if err != nil {
		return err
	}
	return nil

}

// get data from .env : error handling done
func ExtractEnv() (DbConfig, error) {
	dbconfig := DbConfig{
		Username:     os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		Network:      os.Getenv("DB_NETWORK"),
		DatabasePort: os.Getenv("DB_PORT"),
		Host:         os.Getenv("DB_Host"),
	}

	// get the precise unused variables
	values := r.ValueOf(dbconfig)
	keys := r.TypeOf(dbconfig)

	var uninitializedEnvKeys string = ""

	for i := 0; i < values.NumField(); i++ {

		fmt.Println(values.Field(i), " ")

		if values.Field(i).String() == "" {
			uninitializedEnvKeys += (keys.Field(i).Name + ",")
		}

	}

	if len(uninitializedEnvKeys) > 0 {
		return dbconfig, errors.New("Missing ENV Variable : " + uninitializedEnvKeys)
	}

	return dbconfig, nil
}

// @params 1 databasename string
// @return (*sql.DB,error)
func OpenDbConnect(databasename string) (*sql.DB, error) {

	// load .env file
	err := LoadEnv()
	if err != nil {
		log.Fatal("Error: error loading env file")
		return db, err
	}

	// get .env file data
	envdata, err := ExtractEnv()
	if err != nil {
		log.Fatal("Error: error extracting data from env file")
		return db, err
	}

	connectionString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		envdata.Username,
		envdata.Password,
		envdata.Network,
		envdata.Host,
		envdata.DatabasePort,
		databasename,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Print(connectionString)
		return db, err
	}

	// Set parameters for the connection pool
	db.SetMaxOpenConns(10) // Adjust as needed
	db.SetMaxIdleConns(5)  // Adjust as needed

	return db, nil
}
