package mysqlutility

import "log"

// @return (databases,error) ([]string,error)
func ShowDbs() ([]string, error) {

	var databases []string // store and return all dbs in an slice

	db, err := OpenDbConnect("information_schema")
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
