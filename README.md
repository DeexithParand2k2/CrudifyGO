# CrudifyGO
Effortlessly simplify Golang CRUD API development for multiple databases.

## MySQL 

### Table of Contents

* Database Methods
    1. [OpenDbConnect](#opendbconnect)
    2. [PingYourDb](#pingyourdb)
    3. [ListDatabases](#listdatabases)
    4. [CreateDb](#createdb)
    5. [DeleteDb](#deletedb)
    6. [CreateMultipleDb](#createmultipledb) (In-Progress)
    7. [DeleteMultipleDb](#deletemultipledb) (In-Progress)

* Tabular Methods
    1. [ListTablesDb](#listtablesdb)
    2. [GetTableContent](#gettablecontent)

### Database Methods Implementation

<a name="opendbconnect"></a> 

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
    databases, err := mysql.ListDatabases()
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

    err := mysql.CreateDb(database_name)
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

    err := mysql.DeleteDb(database_name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"Message": fmt.Sprintf("Deleted database %s successfully", database_name)})
    ```

### Tabular Methods Implementation

<a name="listtablesdb"></a><br>

1. `ListTablesDb` : List tables in a MySQL database
    + Parameters :
        - databasename (string)
    + Return Type :
        - table ([]string)
        - error
        
    ### Example
    ```go
    database_name := c.Query("databasename")

    tables, err := mysql.ListTablesDb(database_name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "Message":"Listed Tables Successfully",
        "Tables":tables,
    })
    ```

<a name="gettablecontent"></a> 

2. `GetTableContent` : Get all content from a table in DB
    + Parameters :
        - Name of Database (string)
        - Name of Table (string)
        - Empty structure object to store table data (struct obj)
    + Return Type :
        - error
        
    ### Example
    ```
    ```

## PostgreSQL (In-Development)

## SQLite (In-Development)
