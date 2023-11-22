package main

import (
	"fmt"
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

func testShowDbs(c *gin.Context) {

	databases, err := mysqlutility.ShowDbs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{
		"Message":   "Received databases",
		"databases": databases,
	})

}

func testDeleteDb(c *gin.Context) {

	database_name := c.Query("databasename")

	err := mysqlutility.DeleteDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Deleted database %s successfully", database_name)})

}

func testCreateDb(c *gin.Context) {

	database_name := c.Query("databasename")

	err := mysqlutility.CreateDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Created database %s successfully", database_name)})

}

func main() {
	router := gin.Default()

	router.GET("/pingdb", testPingYourDb) // ping db with query as db name
	//router.GET("/gettablesdb",testGetTablesDb) // get tables in a db
	//router.GET("/gettableschema",testGetTableSchema) // get schema of a table

	router.GET("/showdbs", testShowDbs)   // list available dbs
	router.GET("/createdb", testCreateDb) // create a db and ping it
	router.GET("/deletedb", testDeleteDb) // delete existing db and resend present dbs

	router.Run("localhost:8000")
}
