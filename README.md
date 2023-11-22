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
    

## PostgreSQL (In-Development)


## SQLite (In-Development)
