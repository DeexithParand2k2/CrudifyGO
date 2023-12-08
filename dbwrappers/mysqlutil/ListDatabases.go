package mysqlutil

import (
	"log"

	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysqlutil/config"
)

// @return (databases,error) ([]string,error)
func ListDatabases() ([]string, error) {

	var databases []string // store and return all dbs in an slice

	db, err := config.OpenDbConnect("information_schema")
	if err != nil {
		log.Fatal("Error opening database connection:", err)
		return databases, err
	}
	defer db.Close()

	query := "SELECT schema_name FROM schemata"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
		return databases, err
	}
	defer rows.Close()

	for rows.Next() {
		var database_name string

		if err := rows.Scan(&database_name); err != nil {
			log.Fatal("Error scanning rows", err)
		}

		//append in slice
		databases = append(databases, database_name)
	}

	return databases, nil
}
