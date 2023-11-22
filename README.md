# CrudifyGO
Effortlessly simplify Golang CRUD API development for multiple databases

## MySQL 

### Methods

1. `OpenDbConnect` : Establish Connection to a database from database connection pool
    + Parameters :
        - Name of Database (string)
    + Return Type :
        - *sql.DB (pointer to sql.DB object)
        - error
        
    ### Example
    ```
    db,err := OpenDbConnect("databasename")
    if err!=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"Error":err.Error()})
        return
    }
    
2. `PingYourDb` : Check connection to a database
    + Parameters :
        - Name of Database (string)
    + Return Type :
        - connection status
        - error
        
    ### Example
    ```
    pingstatus,err := PingYourDb("databasename")
    if err!=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"Error":err.Error()})
    }
    c.JSON(http.StatusCreated, gin.H{"Message":pingstatus})

3. `ShowDbs` : List all databases on MySQL RDBMS
    + Parameters :
        - Nil
    + Return Type :
        - array of present databases
        - error
        
    ### Example
    ```
    databases, err := mysqlutility.ShowDbs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{
		"Message":   "Received databases",
		"databases": databases,
	})

4. `CreateDb` : Create a database in MySQL
    + Parameters :
        - databasename (string)
    + Return Type :
        - error
        
    ### Example
    ```
    database_name := c.Query("databasename")

	err := mysqlutility.CreateDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Created database %s successfully", database_name)})

5. `DeleteDb` : Delete a database in MySQL
    + Parameters :
        - databasename string
    + Return Type :
        - error
        
    ### Example
    ```
    database_name := c.Query("databasename")

	err := mysqlutility.DeleteDb(database_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Deleted database %s successfully", database_name)})




## PostgreSQL (In-Development)


## SQLite (In-Development)
