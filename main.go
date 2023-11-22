package main

import (
	"net/http"

	mysqlutility "github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql"
	"github.com/gin-gonic/gin"
)

// testing our pingdb func
func testPingYourDb(c *gin.Context) {

	databaseNameQuery := c.Query("databasename")
	status, err := mysqlutility.PingYourDb(databaseNameQuery)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"Message": status})
}

func main() {
	router := gin.Default()

	router.GET("/pingdb", testPingYourDb) // ping db with query as db name
	//router.GET("/createdb",CreateYourDb) // create a db and ping it
	//router.GET("/gettablesdb",GetTablesDb) // get tables in a db
	//router.GET("/gettableschema",GetTableSchema) // get schema of a table

	router.Run("localhost:8000")
}
