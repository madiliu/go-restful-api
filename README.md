# user_assignment
* A simple RESTful API assigment base on Go Gin, Postgresql, and Docker
### Environment
* Go 1.17
* Go Gin 1.8.1
* Postgresql 14.5
* Docker 20.10.12

## How to use this
### 1. clone this repo
```

```
### 2. docker compose
### 3. run the server on port 8080
### 4. check server is running


## API Document
### [1] Get User List
#### `GET: /user`
#### Description: obtain the user list
#### Parameters: none
#### Response Body
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

### [2] Get User
#### `POST: /user/:user_id`
#### Description: query the user information by inputting id
#### Parameters: 
| No  | Param   | Description | Data Type | Required  | Example |
|---- |---------|----------|-----------|----------  |---------|
| 1   | user_id | 流水號      | int32     | Y                   | 1       |

#### Response Body
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



### [3] Create User
#### `POST: /user`
#### Description: create single user information and return the inserted data
#### Parameters
| No  | Param | Description | Data Type | Required      | Example |
|---- |-------|-------------|-----------|----------  |--------|
| 1   | name    | 名稱          | string    | Y                     | korr   |

#### Response Body
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

### [4] Update User
#### `PUT: /user/:user_id`
#### Description: update single user information and return the updated data
#### Parameters
| No  | Param | Description | Data Type | Required | Example |
|---- |-------|--------|-----------|----------|---------|
| 1   | id    | 流水號    | int32     | Y        | 5       |
| 2   | name  | 名稱     | string    | Y        | sokka   |

#### Response Body
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

### [5] Delete
#### `DELETE: /user/:user_id`
#### Description: delete single user information by providing user id
#### Parameters
| No  | Param   | Description | Data Type | Required    | Example |
|---- |---------|-------------|-----------|---------- |---------|
| 1   | user_id | 流水號         | int32     | Y                    | 6       |

#### Response Body
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


