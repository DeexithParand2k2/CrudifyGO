package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	mysql "github.com/DeexithParand2k2/CrudifyGO/dbwrappers/mysql"
)

// database methods

func testPingYourDb(c *gin.Context) {

	databaseNameQuery := c.Query("databasename")
	status, err := mysql.PingYourDb(databaseNameQuery)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"Message": status})
}

func testListDatabases(c *gin.Context) {

	databases, err := mysql.ListDatabases()
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

	err := mysql.DeleteDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Deleted database %s successfully", database_name)})

}

func testCreateDb(c *gin.Context) {

	database_name := c.Query("databasename")

	err := mysql.CreateDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Created database %s successfully", database_name)})

}

// tabular Methods

func testListTablesDb(c *gin.Context) {
	database_name := c.Query("databasename")

	tables, err := mysql.ListTablesDb(database_name)
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

type customer struct {
	Cid   int    `json:"cid"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Total int    `json:"total"`
}

/*
* 2 queries available
* pass the object of table content you need
 */
func testGetTableContent(c *gin.Context) {

	database_name := c.Query("databasename")
	table_name := c.Query("tablename")

	var custTable []customer

	err := mysql.GetTableContent(database_name, table_name, &custTable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"Message":       "Returned table content",
		"Table Content": custTable,
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
