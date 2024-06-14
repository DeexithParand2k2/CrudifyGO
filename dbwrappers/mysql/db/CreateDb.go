package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql/config"
)

// @params databasename (string)
// @return query execution status (error)
func CreateDb(databasename string) error {

	if databasename == "" {
		return errors.New("query doesn't contain any database name")
	}

	db, err := config.OpenDbConnect("") // connection string stops with /
	if err != nil {
		log.Fatal("opening db error", err.Error())
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("create database %s", databasename))
	if err != nil {
		log.Fatal("Error in executing query", err.Error())
		return err
	}

	return nil
}
