package mysqlutil

import "github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysqlutil/config"

// @params 1 databasename string
// @return (string,error)
func PingYourDb(databaseName string) (string, error) {
	db, err := config.OpenDbConnect(databaseName) // returns
	if err != nil {
		return "", err
	}
	defer db.Close() // after pinging close it

	if err = db.Ping(); err != nil {
		return "", err
	}

	return "Success pinging db" + databaseName, nil
}
