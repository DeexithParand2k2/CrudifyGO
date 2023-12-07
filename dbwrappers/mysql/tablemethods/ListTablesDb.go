package mysqlutil

import (
	"log"

	mysqlutil "github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql"
)

func ListTablesDb(databasename string) ([]string, error) {

	var tables []string

	db, err := mysqlutil.OpenDbConnect(databasename)
	if err != nil {
		log.Print("Error opening db", err)
		return tables, err
	}
	defer db.Close()

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Print("Error executing query", err)
		return tables, err
	}
	defer rows.Close()

	var tablename string
	for rows.Next() {

		if err = rows.Scan(&tablename); err != nil {
			log.Fatal("Error reading tablename rows")
		}

		tables = append(tables, tablename)

	}

	return tables, nil
}
