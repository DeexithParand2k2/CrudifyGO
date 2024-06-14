package table

import (
	"fmt"
	"log"
	r "reflect"

	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql/config"
)

/*
* @param databasename (string)
* @param tablename (string)
* @return (error)
 */
func GetTableContent(databasename string, tablename string, tableStore interface{}) error {

	db, err := config.OpenDbConnect(databasename)
	if err != nil {
		log.Print("Error opening db")
		return err
	}
	defer db.Close()

	// get content based on query
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tablename))
	if err != nil {
		log.Print("Error retrieving data from db")
		return err
	}
	defer rows.Close()

	tableType := r.TypeOf(tableStore).Elem().Elem() // type of passed array

	tableSlice := r.MakeSlice(r.SliceOf(tableType), 0, 0) // slice to store obj of this type

	for rows.Next() {

		tableObj := r.New(tableType).Elem()

		fields := make([]interface{}, tableObj.NumField())
		for i := 0; i < tableObj.NumField(); i++ {
			field := tableObj.Field(i)
			if field.CanInterface() {
				fields[i] = field.Addr().Interface()
			}
		}

		if err := rows.Scan(fields...); err != nil {
			return err
		}

		tableSlice = r.Append(tableSlice, tableObj)

	}

	fmt.Println("here is the slice", tableSlice)

	if r.ValueOf(tableStore).Kind() == r.Ptr {
		r.ValueOf(tableStore).Elem().Set(tableSlice)
	} else {
		return fmt.Errorf("tableStore must be a pointer to a slice")
	}

	return nil
}
