package mysql

import (
	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql/db"
	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql/table"
)

func ListTablesDb(databasename string) ([]string, error) {
	return table.ListTablesDb(databasename)
}

func CreateDb(databasename string) error {
	return db.CreateDb(databasename)
}

func DeleteDb(databasename string) error {
	return db.DeleteDb(databasename)
}

func GetTableContent(databasename string, tablename string, tableStore interface{}) error {
	return table.GetTableContent(databasename, tablename, tableStore)
}
