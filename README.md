# CrudifyGO
Effortlessly simplify Golang CRUD API development for multiple databases

## MySQL 

### Table of Contents

* Database Methods
    1. [OpenDbConnect](#opendbconnect)
    2. [PingYourDb](#pingyourdb)
    3. [ListDatabases](#listdatabases)
    4. [CreateDb](#createdb)
    5. [DeleteDb](#deletedb)
* Tabular Methods
    1. [ListDatabases](#listdatabases)

### Methods Implementation

<a name="opendbconnect"></a><br>
1. `OpenDbConnect` : Establish Connection to a database from database connection pool
    + Parameters :
        - Name of Database (string)
    + Return Type :
        - *sql.DB (pointer to sql.DB object)
        - error
        
    ### Example
    ```go
    db,err := OpenDbConnect("databasename")
    if err!=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"Error":err.Error()})
        return
    }
    ```

<a name="pingyourdb"></a><br>
2. `PingYourDb` : Check connection to a database
    + Parameters :
        - Name of Database (string)
    + Return Type :
        - connection status
        - error
        
    ### Example
    ```go
    pingstatus,err := PingYourDb("databasename")
    if err!=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"Error":err.Error()})
    }
    c.JSON(http.StatusCreated, gin.H{"Message":pingstatus})
    ```

<a name="listdatabases"></a><br>
3. `ListDatabases` : List all databases on MySQL RDBMS
    + Parameters :
        - Nil
    + Return Type :
        - array of present databases
        - error
        
    ### Example
    ```go
    databases, err := mysqlutility.ListDatabases()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusAccepted, gin.H{
        "Message":   "Received databases",
        "databases": databases,
    })
    ```

<a name="createdb"></a><br>
4. `CreateDb` : Create a database in MySQL
    + Parameters :
        - databasename (string)
    + Return Type :
        - error
        
    ### Example
    ```go
    database_name := c.Query("databasename")

    err := mysqlutility.CreateDb(database_name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Created database %s successfully", database_name)})
    ```

<a name="deletedb"></a><br>
5. `DeleteDb` : Delete a database in MySQL
    + Parameters :
        - databasename string
    + Return Type :
        - error
        
    ### Example
    ```go
    database_name := c.Query("databasename")

    err := mysqlutility.DeleteDb(database_name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Deleted database %s successfully", database_name)})
    ```

## PostgreSQL (In-Development)

## SQLite (In-Development)
