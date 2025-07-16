# 📄 Task Management REST API Documentation

This API allows **creating, retrieving, updating, and deleting tasks** stored in **in-memory storage using Go and Gin Framework**.

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
Fetch and display all tasks currently in the in-memory store.

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks`
-   **Headers:** `Content-Type: application/json; charset=utf-8`

**Response:**

-   **Status Code:** 200 OK
-   **Example JSON:**

```json
{
    "Tasks": [
        {
            "id": "1",
            "title": "Task 4",
            "description": "Backend with Go",
            "status": "In progress",
            "due_date": "2025-07-21T14:24:14.2808923+03:00"
        },
        {
            "id": "2",
            "title": "Design API Endpoints",
            "description": "Plan REST API structure and routes for task management",
            "status": "Pending",
            "due_date": "2025-07-18T14:24:14.2808923+03:00"
        },
        {
            "id": "3",
            "title": "Setup Database",
            "description": "Initialize PostgreSQL and create necessary tables",
            "status": "Completed",
            "due_date": "2025-07-13T14:24:14.2808923+03:00"
        }
    ]
}
```

**Errors:**  
None (returns an empty slice if there are no tasks).

---

### 2️⃣ GET `/tasks/:id`

**Description:**  
Fetch and display a single task by its ID.

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks/:id`
-   **Example:** `http://localhost:8080/tasks/1`

**Response:**

-   **Status Code:** 200 OK
-   **Example JSON:**

```json
{
    "Task": {
        "id": "1",
        "title": "Task 4",
        "description": "Backend with Go",
        "status": "In progress",
        "due_date": "2025-07-21T14:24:14.2808923+03:00"
    }
}
```

**Errors:**

-   `404 Not Found` if the task is not found.

```json
{
    "message": "task not found"
}
```

---

### 3️⃣ PUT `/tasks/:id`

**Description:**  
Update a specific task by its ID with new details.

**Request:**

-   **Method:** PUT
-   **URL:** `http://localhost:8080/tasks/:id`
-   **Headers:** `Content-Type: application/json`
-   **Body Example:**

```json
{
    "title": "Task 4 Updated",
    "description": "Updated backend with Go and Gin",
    "status": "Completed",
    "dueDate": "2025-07-21T12:00:00Z"
}
```

**Response:**

-   **Status Code:** 200 OK
-   **Example JSON:**

```json
{
    "message": "Task updated Successfully!"
}
```

**Errors:**

-   `400 Bad Request` if the request payload is malformed.

```json
{
    "error": "Invalid request payload"
}
```

-   `404 Not Found` if the task is not found.

```json
{
    "message": "task not found"
}
```

---

### 4️⃣ POST `/tasks`

**Description:**  
Create a new task.

**Request:**

-   **Method:** POST
-   **URL:** `http://localhost:8080/tasks`
-   **Headers:** `Content-Type: application/json`
-   **Body Example:**

```json
{
    "title": "Write Documentation",
    "description": "Create API documentation using Markdown",
    "status": "Pending",
    "dueDate": "2025-07-25T23:59:59Z"
}
```

**Response:**

-   **Status Code:** 201 Created
-   **Example JSON:**

```json
{
    "id": "4",
    "title": "Write Documentation",
    "description": "Create API documentation using Markdown",
    "status": "Pending",
    "dueDate": "2025-07-25T23:59:59Z"
}
```

**Errors:**

-   `400 Bad Request` if JSON is malformed or required fields are missing.

```json
{
    "error": "Invalid request payload"
}
```

---

### 5️⃣ DELETE `/tasks/:id`

**Description:**  
Delete a task by its ID.

**Request:**

-   **Method:** DELETE
-   **URL:** `http://localhost:8080/tasks/:id`
-   **Example:** `http://localhost:8080/tasks/1`

**Response:**

-   **Status Code:** 200 OK
-   **Example JSON:**

```json
{
    "message": "Task deleted!"
}
```

**Errors:**

-   `404 Not Found` if the task is not found.

```json
{
    "message": "task not found"
}
```

---

## 🧪 Testing

-   All endpoints were tested using **Postman (VS Code extension)**.
-   Requests return the expected responses and error codes as documented.

---

## 🚀 Running the API

1. Ensure Go is installed.
2. Run:
    ```
    go mod tidy
    go run main.go
    ```
3. Server will run at `http://localhost:8080`.
4. Use Postman, Curl, or VS Code REST Client to interact with the API using the documented endpoints.

---

## ℹ️ Notes

-   Data is stored in-memory and resets on server restart.
-   No authentication is implemented in this phase.
-   Ensure `Content-Type: application/json` for POST and PUT requests.

---
