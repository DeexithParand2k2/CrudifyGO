package main

import (
	"fmt"
	"net/http"

	mysqlutil "github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql"
	"github.com/gin-gonic/gin"
)

// database methods

func testPingYourDb(c *gin.Context) {

	databaseNameQuery := c.Query("databasename")
	status, err := mysqlutil.PingYourDb(databaseNameQuery)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"Message": status})
}

func testListDatabases(c *gin.Context) {

	databases, err := mysqlutil.ListDatabases()
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

	err := mysqlutil.DeleteDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Deleted database %s successfully", database_name)})

}

func testCreateDb(c *gin.Context) {

	database_name := c.Query("databasename")

	err := mysqlutil.CreateDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Created database %s successfully", database_name)})

}

// tabular Methods

func testListTablesDb(c *gin.Context) {
	database_name := c.Query("databasename")

	tables, err := mysqlutil.ListTablesDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if len(tables) == 0 {
		c.JSON(http.StatusCreated, gin.H{
			"Message": "Listed Tables Successfully",
			"Tables":  "Empty",
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"Message": "Listed Tables Successfully",
			"Tables":  tables,
		})
	}
}

// 2 queries available
func testGetTableContent(c *gin.Context) {

	database_name := c.Query("databasename")
	table_name := c.Query("tablename")

	data, err := mysqlutil.GetTableContent(database_name, table_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"Message":         "Returned table content",
		"Tabular Content": data,
	})

}

func main() {
	router := gin.Default()

	// db operations

	router.GET("/pingdb", testPingYourDb)     // ping db with query as db name
	router.GET("/listdbs", testListDatabases) // list available dbs
	router.GET("/createdb", testCreateDb)     // create a db and ping it
	router.GET("/deletedb", testDeleteDb)     // delete existing db and resend present dbs

	// table operations

	router.GET("/listtablesdb", testListTablesDb)       // get tables in a db
	router.GET("/gettablecontent", testGetTableContent) // get table content

	router.Run("localhost:8000")
}
