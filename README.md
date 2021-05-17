# Bookstore - Users API

This project is one of the services built during a Golang course.  
[How to design & develop REST microservices in Golang (Go)](https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go).  
  
The service projects and API do create, read, update and delete users.    
It uses [Gin Gonic](https://github.com/gin-gonic/gin) as web framework and stores users in a MySQL database.  

## Running

Make sure you have Go, Docker and Docker Compose installed to facilitate the service execution.

1. Start MySQL Docker container: `docker-compose up -d`
2. Download service dependencies: `go get ./...`
3. Run service: `go run main.go`
4. Access the service endpoints at `http://localhost:8080`

## Endpoints

The file `Bookstore - Users API.postman_collection.json` has a collection of the service endpoints that can be import on Postman.

### `POST /users`
**Description:** Create new user  
**Body:**
```json
{
  "first_name": "Leandro",
  "last_name": "Coutinho",
  "email": "leandro.coutinho@email.com",
  "status":"active",
  "password":"12345"
}
```
**Return Status:**
* **201 Created:** User created successfully
* **400 Bad Request:** Invalid request or validation error
* **500 Internal Server Error:** Unexpected error

### `GET /users/{user_id}`
**Description:** Get user by id  
**Headers:** 
* X-Public = true/false (Return all data or only public data)  

**Return Status:**
* **200 OK:** User returned
* **400 Bad Request:** Invalid request or validation error
* **500 Internal Server Error:** Unexpected error

### `PUT /users/{user_id}`
**Description:** Complete user update
**Body:**
```json
{
  "first_name": "Leandro",
  "last_name": "Coutinho",
  "email": "leandro.coutinho@email.com",
  "status":"active"
}
```
**Headers:**
* X-Public = true/false (Return all data or only public data)

**Return Status:**
* **200 OK:** User updated successfully
* **400 Bad Request:** Invalid request or validation error
* **500 Internal Server Error:** Unexpected error

### `PATCH /users/{user_id}`
**Description:** Partial user update
**Body:**
```json
{
  "email": "leandro@new.com",
  "status":"inative"
}
```
**Headers:**
* X-Public = true/false (Return all data or only public data)

**Return Status:**
* **200 OK:** User updated successfully
* **400 Bad Request:** Invalid request or validation error
* **500 Internal Server Error:** Unexpected error

### `DELETE /users/{user_id}`
**Description:** Delete user  
**Return Status:**
* **204 No Content:** User deleted successfully
* **400 Bad Request:** Invalid request or validation error
* **500 Internal Server Error:** Unexpected error

### `GET /internal/users/search?status={user_status}`
**Description:** Search users by status  
**Headers:**
* X-Public = true/false (Return all data or only public data)

**Return Status:**
* **200 OK: ** Users found. Returns empty list if no users found.
* **400 Bad Request:** Invalid request or validation error
* **500 Internal Server Error:** Unexpected error
