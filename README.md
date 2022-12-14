<div><img src ="https://user-images.githubusercontent.com/90917906/196348724-a25be68c-2444-4399-b8f9-096df309b996.png"></div>

# Go RESTful APIs
RESTful APIs built with Gin-Gonic, PostgreSQL, and Docker<br>
 [Installation and Execution](https://github.com/madiliu/go-restful-api/blob/main/README.md#-installation-and-execution) 
 |  [API Documentation](https://github.com/madiliu/go-restful-api/blob/main/README.md#-api-documentation)

##  Stacks
* Go 1.17
* Go Gin 1.8.1
* PostgreSQL 14.5
* Docker 20.10.12

##  Features
* Restful APIs to read/write
* Use relational database for storage
* Use JSON format for communication
* API documentation (README)
* Unit tests with httptest and testify
* Logrus to record API results
* Utilize docker compose for two services

##  Installation and Execution
### 1. clone this repo
```
$ git clone https://github.com/madiliu/user-assignment.git
$ cd user-assignment
```
### 2. replace environment variables in config.env with yours
```
POSTGRES_USER={your_user}
POSTGRES_DB={your_db}
POSTGRES_PASSWORD={your_password}
```
### 3. build docker-compose
```
$ docker-compose build
```
### 4. run docker-compose
```
$ docker-compose up
```
### 5. check server is running
```
$ curl http://localhost:8080/user
```

## API Documentation
1. [`GET: /user`](https://github.com/madiliu/go-restful-api/blob/main/README.md#1-get-user)
2. [`GET: /user/:user_id`](https://github.com/madiliu/go-restful-api/blob/main/README.md#2-get-useruser_id)
3. [`POST: /user`](https://github.com/madiliu/go-restful-api/blob/main/README.md#3-post-user)
4. [`PUT: /user/:user_id`](https://github.com/madiliu/go-restful-api/blob/main/README.md#4-put-useruser_id)
5. [`DELETE: /user/:user_id`](https://github.com/madiliu/go-restful-api/blob/main/README.md#5-delete-useruser_id)

### 1. `GET: /user`
* Description: obtain the user list
* Parameters: none
* Response Body
  ##### 200 OK
  ```json
  [
      {
          "ID": 1,
          "Name": "peter"
      },
      {
          "ID": 2,
          "Name": "amy"
      },
      {
          "ID": 3,
          "Name": "bob"
      },
      {
          "ID": 4,
          "Name": "catherine"
      },
      {
          "ID": 5,
          "Name": "david"
      }
  ]
  ```
  ##### 404 Not Found
  ```json
  {
      "message": "Users information unavailable"
  }
  ```
  #### 500 Internal Server Error
  ```json
  {
      "error": "xxxx",
      "message": "Please contact Madelain for further assistance"
  }
  ```

### 2. `GET: /user/:user_id`
* Description: query the user information by inputting id
* Parameters 

  | No  | Param   | Description | Data Type | Required  | Example |
  |---- |---------|----------|-----------|----------  |---------|
  | 1   | user_id | ?????????      | int32     | Y        | 1       |
 
* Response Body
  ##### 200 OK
  ```json
  {
      "user_id": 1,
      "name": "peter"
  }
  ```
  ##### 400 Bad Request
  ```json
  {
      "error": "strconv.ParseInt: parsing \"1jdw\": invalid syntax",
      "message": "Please provide id in the correct format"
  }
  ```
  ##### 404 Not Found
  ```json
  {
      "error": "sql: no rows in result set",
      "message": "The inputted id does not exist"
  }
  ```
  #### 500 Internal Server Error
  ```json
  {
      "error": "xxxx",
      "message": "Please contact Madelain for further assistance"
  }
  ```

### 3. `POST: /user`
* Description: create single user information and return the inserted data
* Parameters
  | No  | Param | Description | Data Type | Required      | Example |
  |---- |-------|-------------|-----------|----------  |--------|
  | 1   | name    | ??????          | string    | Y         | korr   |

* Response Body
  ##### 200 OK
  ```json
  {
      "user_id": 6,
      "name": "korra"
  }
  ```
  ##### 400 Bad Request
  ```json
  {
      "error": "Key: 'apiUser.Name' Error:Field validation for 'Name' failed on the 'max' tag",
      "message": "Please provide name within 20 alphabetical characters"
  }
  ```
  ##### 500 Internal Server Error
  ```json
  {
      "data": "Failed to insert the currency, error details: ...."
  }
  ```

### 4. `PUT: /user/:user_id`
* Description: update single user information and return the updated data
* Parameters
  | No  | Param | Description | Data Type | Required | Example |
  |---- |-------|--------|-----------|----------|---------|
  | 1   | user_id    | ?????????    | int32     | Y        | 5       |
  | 2   | name  | ??????     | string    | Y        | sokka   |

* Response Body
  ##### 200 OK
  ```json
  {
      "user_id": 5,
      "name": "sokka"
  }
  ```
  #### 400 Bad Request
  ```json
  {
      "error": "Key: 'apiUser.Name' Error:Field validation for 'Name' failed on the 'max' tag",
      "message": "Please provide name within 20 alphabetical characters"
  }
  ```
  ##### 404 Not Found
  ```json
  {
      "error": "sql: no rows in result set",
      "message": "The inputted id does not exist"
  }
  ```
  #### 500 Internal Server Error
  ```json
  {
      "error": "xxxx",
      "message": "Please contact Madelain for further assistance"
  }
  ```

### 5. `DELETE: /user/:user_id`
* Description: delete single user information by providing user id
* Parameters
  | No  | Param   | Description | Data Type | Required    | Example |
  |---- |---------|-------------|-----------|---------- |---------|
  | 1   | user_id | ?????????         | int32     | Y          | 6       |

* Response Body
  ##### 200 OK
  ```json
  {
      "data": "Deleted successfully"
  }
  ```
  ##### 400 Bad Request
  ```json
  {
      "error": "strconv.ParseInt: parsing \"dwdwdw\": invalid syntax",
      "message": "Please provide id in the correct format"
  }
  ```
  ##### 404 Not Found
  ```json
  {
      "error": "sql: no rows in result set",
      "message": "The inputted id does not exist"
  }
  ```
  #### 500 Internal Server Error
  ```json
  {
      "error": "xxxx",
      "message": "Please contact Madelain for further assistance"
  }
  ```


