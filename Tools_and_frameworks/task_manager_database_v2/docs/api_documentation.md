# 📄 Task Management REST API Documentation

This API allows **creating, retrieving, updating, and deleting tasks**.  
It was initially backed by an **in-memory store**, but now uses **MongoDB** as the primary data store.

---

## Base URL

```
http://localhost:8080
```

---

## Endpoints

---

### 1️⃣ GET `/tasks`

**Description:**  
Fetch and display all tasks currently in the MongoDB store.

**Request:**

- **Method:** GET  
- **URL:** `http://localhost:8080/tasks`  
- **Headers:** `Content-Type: application/json; charset=utf-8`

**Response:**

- **Status Code:** 200 OK  
- **Example JSON:**

```json
{
  "Tasks": [
    {
      "_id": "64e4f1d9bc19d2e91f6dcaa1",
      "title": "Task 4",
      "description": "Backend with Go",
      "status": "In progress",
      "due_date": "2025-07-21T12:00:00Z"
    }
  ]
}
```

**Errors:**  
Returns an empty array if there are no tasks.

---

### 2️⃣ GET `/tasks/:id`

**Description:**  
Fetch and display a single task by its MongoDB ObjectID.

**Request:**

- **Method:** GET  
- **URL:** `http://localhost:8080/tasks/:id`  
- **Example:** `http://localhost:8080/tasks/64e4f1d9bc19d2e91f6dcaa1`

**Response:**

- **Status Code:** 200 OK  
- **Example JSON:**

```json
{
  "Task": {
    "_id": "64e4f1d9bc19d2e91f6dcaa1",
    "title": "Task 4",
    "description": "Backend with Go",
    "status": "In progress",
    "due_date": "2025-07-21T12:00:00Z"
  }
}
```

**Errors:**

- `404 Not Found` if the task is not found.

```json
{
  "message": "task not found"
}
```
- `500 Internal Sever Error` if there is a database handling error 
```json 
{
  "error": "Database failure"
}
```
---

### 3️⃣ PUT `/tasks/:id`

**Description:**  
Update a specific task by its ID with new details.

**Request:**

- **Method:** PUT  
- **URL:** `http://localhost:8080/tasks/:id`  
- **Headers:** `Content-Type: application/json`  
- **Body Example:**

```json
{
  "title": "Task 4 Updated",
  "description": "Updated backend with Go and Gin",
  "status": "Completed",
  "dueDate": "2025-07-21T12:00:00Z"
}
```

**Response:**

- **Status Code:** 200 OK  
- **Example JSON:**

```json
{
  "message": "Task updated Successfully!"
}
```

**Errors:**

- `400 Bad Request` if the request payload is malformed.

```json
{
  "error": "Invalid request payload"
}
```

- `404 Not Found` if the task is not found.

```json
{
  "message": "task not found"
}
```
- `500 Internal Server Error` if there is a database handling error 
```json 
{
  "error": "Database failure"
}
```
---

### 4️⃣ POST `/tasks`

**Description:**  
Create a new task.

**Request:**

- **Method:** POST  
- **URL:** `http://localhost:8080/tasks`  
- **Headers:** `Content-Type: application/json`  
- **Body Example:**

```json
{
  "title": "Write Documentation",
  "description": "Create API documentation using Markdown",
  "status": "Pending",
  "dueDate": "2025-07-25T23:59:59Z"
}
```

**Response:**

- **Status Code:** 201 Created  
- **Example JSON:**

```json
{
  "_id": "64e4f2b0bc19d2e91f6dca11",
  "title": "Write Documentation",
  "description": "Create API documentation using Markdown",
  "status": "Pending",
  "dueDate": "2025-07-25T23:59:59Z"
}
```

**Errors:**

- `400 Bad Request` if JSON is malformed or required fields are missing.

```json
{
  "error": "Invalid request payload"
}
```
- `500 Internal server error` if there is a database handling error 
```json 
{
  "error": "Database failure"
}
```
---

### 5️⃣ DELETE `/tasks/:id`

**Description:**  
Delete a task by its ID.

**Request:**

- **Method:** DELETE  
- **URL:** `http://localhost:8080/tasks/:id`  
- **Example:** `http://localhost:8080/tasks/64e4f1d9bc19d2e91f6dcaa1`

**Response:**

- **Status Code:** 200 OK  
- **Example JSON:**

```json
{
  "message": "Task deleted!"
}
```

**Errors:**

- `404 Not Found` if the task is not found.

```json
{
  "message": "task not found"
}
```

---

## 🧪 Testing

- All endpoints were tested using **Postman (VS Code extension)**.
- Requests return the expected responses and error codes as documented.

---

## 🚀 Running the API

1. Ensure Go is installed.  
2. Run:
    ```bash
    go mod tidy
    go run main.go
    ```
3. Server will run at `http://localhost:8080`.  
4. Use Postman, Curl, or VS Code REST Client to interact with the API using the documented endpoints.

---

## 🧩 MongoDB Integration Process

### ✅ Prerequisites

- MongoDB installed and running locally or hosted (e.g. MongoDB Atlas).  
- Go MongoDB driver installed:

```bash
go get go.mongodb.org/mongo-driver/mongo
```

- `.env` example:

```
MONGODB_URI=mongodb://localhost:27017
DB_NAME=taskmanager
```

### ✅ MongoDB Connection (Go)

```go
import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "os"
  "time"
)

var taskCollection *mongo.Collection

func InitDB() {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
  if err != nil {
    log.Fatal(err)
  }

  taskCollection = client.Database(os.Getenv("DB_NAME")).Collection("tasks")
}
```

---

## 📘 API Documentation Update Summary

- Previously used in-memory slices, now stores tasks in the `tasks` collection in MongoDB.
- Error handling and status codes updated to include database handling errors.

---

## ℹ️ Notes

- MongoDB replaces in-memory data and persists across restarts.
- No authentication is implemented in this phase.
- Ensure `Content-Type: application/json` for POST and PUT requests.

---

Here is an additional documentation made using Postman: [Link](https://documenter.getpostman.com/view/46777269/2sB34ijKKU) / Make sure to run locally to see the results.
