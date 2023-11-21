package main 

import (
	"fmt"
	"os"
	"errors"
	r "reflect"
	"net/http"
	"github.com/gin-gonic/gin"
	godot "github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

// from env
type DbConfig struct {
	Username  string
	Password  string
	Network   string
	DatabasePort   string // Default PORT for MySQL
	Host   	  string // Localhost IP Addr
	DatabaseName   string // Our DB
}


// load local .env : error handling done
func LoadEnv() error {
	
	err := godot.Load(".env")
	if err!=nil{
		return err 
	}
	return nil
}

// get data from .env : error handling done
func ExtractEnv() (DbConfig,error){
	dbconfig := DbConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Network: os.Getenv("DB_NETWORK"),
		DatabasePort: os.Getenv("DB_PORT"),
		Host: os.Getenv("DB_Host"),
		DatabaseName: os.Getenv("DB_NAME"),
	}

	// get the precise unused variables
	values := r.ValueOf(dbconfig)
	keys := r.TypeOf(dbconfig)

	var uninitializedEnvKeys string = ""

	for i:=0; i<values.NumField(); i++ {

		fmt.Println(values.Field(i)," ")

		if values.Field(i).String()=="" {
			uninitializedEnvKeys += (keys.Field(i).Name+",")
		}

	}

	if len(uninitializedEnvKeys)>0{
		return dbconfig, errors.New("Missing ENV Variable : "+uninitializedEnvKeys)
	}

	return dbconfig,nil
}

// @ query key : "databasename"
func PingYourDb(c *gin.Context){

	queryDbName := c.Query("databasename")

	if queryDbName==""{
		log.Print("Error: Query doesn't contain any database name")
		c.JSON(http.StatusBadRequest, gin.H{"Query Error": "Query doesn't contain any database name"})
		return
	}

	// load .env file
	err := LoadEnv()
	if err!=nil{
		log.Print("Error: error loading env file")
		c.JSON(http.StatusInternalServerError, gin.H{"Loading Error": err.Error()})
		return 
	}

	// get .env file data
	envdata,err := ExtractEnv()
	if err!=nil{
		log.Print("Error: error extracting data from env file")
		c.JSON(http.StatusInternalServerError, gin.H{"Extraction Error": err.Error()})
		return 
	}

	// extracted data from .env file - now connectdb
	var connectionString string = envdata.Username+":"+envdata.Password+"@"+envdata.Network+"("+envdata.Host+":"+envdata.DatabasePort+")/"+queryDbName
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Print("Error: error on connection to db")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	defer db.Close()

	if err = db.Ping(); err != nil{
		log.Print("Error: Pinging error to db")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusInternalServerError)
		return	
	}
 
	c.IndentedJSON(http.StatusCreated, gin.H{"Success":"Connection to DB "+queryDbName+" successfull"})

}

func main(){
	router := gin.Default()

	router.GET("/pingdb",PingYourDb) // ping db with query as db name

	router.Run("localhost:8000")
}

