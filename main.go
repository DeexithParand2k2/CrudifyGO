package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql"
)

func main(){
	router := gin.Default()

	router.GET("/pingdb",mysqlutility.PingYourDb) // ping db with query as db name
	//router.GET("/createdb",CreateYourDb) // create a db and ping it
	//router.GET("/gettablesdb",GetTablesDb) // get tables in a db
	//router.GET("/gettableschema",GetTableSchema) // get schema of a table

	router.Run("localhost:8000")
}

