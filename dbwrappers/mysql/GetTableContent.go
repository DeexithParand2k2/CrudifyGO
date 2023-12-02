package mysqlutil

import (
	"fmt"
	"log"
)

/*
* @param databasename (string)
* @param tablename (string)
* @return tablecontent ([][]string)
 */
func GetTableContent(databasename string, tablename string) ([][]string, error) {

	db, err := OpenDbConnect(databasename)
	if err != nil {
		log.Print("Error opening db")
		return [][]string{}, err
	}
	defer db.Close()

	// check if connection to table is live
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tablename))
	if err != nil {
		log.Print("Error retrieving data from db")
		return [][]string{}, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Print("Not able to get columns")
		return [][]string{}, err
	}

	for index, val := range columns {
		fmt.Println("Columns", index, ":", val)
	}

	fmt.Println("Length", len(columns))

	//fmt.Println("Output #######", rows)

	return [][]string{}, nil
}
