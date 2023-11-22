package mysqlutility

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @ query key : "databasename"
func PingYourDb(c *gin.Context){

	queryDbName := c.Query("databasename")

	if queryDbName==""{
		log.Fatal("Error: Query doesn't contain any database name")
		c.JSON(http.StatusBadRequest, gin.H{"Query Error": "Query doesn't contain any database name"})
		return
	}
	
	db, err := OpenDbConnect(queryDbName)
	if err != nil {
		log.Print("Error: error on connection to db")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	defer db.Close()

	if err = db.Ping(); err != nil{
		log.Print("Error: Pinging error to db")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return	
	}
 
	c.IndentedJSON(http.StatusCreated, gin.H{"Success":"Connection to DB "+queryDbName+" successfull"})

}
