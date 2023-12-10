package mysqlutil

import (
	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysqlutil/db"
	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysqlutil/table"
)

func ListTablesDb(databasename string) ([]string, error) {
	return table.ListTablesDb(databasename)
}

func CreateDb(databasename string) error {
	return db.CreateDb(databasename)
}