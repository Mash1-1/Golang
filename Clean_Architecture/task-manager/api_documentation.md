# 📄 Task Management REST API Documentation

This API allows **creating, retrieving, updating, and deleting tasks**.  
It was initially backed by an **in-memory store**, but now uses **MongoDB** as the primary data store and follows a **Clean Architecture** structure for better separation of concerns.

---

## 🧱 Clean Architecture Overview

This project adopts **Clean Architecture** to promote separation of concerns, testability, and scalability. Each layer has a clearly defined responsibility.

### 📁 Project Folder Structure

```
task-manager/
├── Delivery/
│   ├── main.go                       # Application entry point
│   ├── controllers/                 # HTTP handlers/controllers
│   │   └── controller.go
│   └── routers/                     # HTTP route setup
│       └── router.go
├── Domain/                          # Core business models & interfaces
│   └── domain.go
├── Infrastructure/                 # External service implementations
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/                   # Data persistence logic (MongoDB)
│   ├── task_repository.go
│   └── user_repository.go
└── Usecases/                       # Application-specific business logic
    ├── task_usecases.go
    └── user_usecases.go
```

---

### 🔄 Layer Responsibilities

#### 1. **Domain Layer**

-   **Path:** `Domain/`
-   **Purpose:** Defines core business models and interfaces (e.g., Task, User, TaskRepository interface).
-   **Contains:** Pure Go structs, no external dependencies.

#### 2. **Usecase Layer**

-   **Path:** `Usecases/`
-   **Purpose:** Encapsulates application-specific business logic using domain models.
-   **Example:** `CreateTask`, `GetTasksByUserID`, `LoginUser`, etc.
-   **Depends on:** Domain interfaces, not infrastructure or frameworks.

#### 3. **Repository Layer**

-   **Path:** `Repositories/`
-   **Purpose:** Implements domain interfaces using MongoDB or any external DB.
-   **Example:** `task_repository.go` implements `TaskRepository` interface.
-   **Depends on:** MongoDB driver (external), but not Gin or Delivery layer.

#### 4. **Infrastructure Layer**

-   **Path:** `Infrastructure/`
-   **Purpose:** Provides JWT handling, password hashing, and middleware logic.
-   **Example:** `jwt_service.go`, `auth_middleWare.go`

#### 5. **Delivery Layer**

-   **Path:** `Delivery/`
-   **Purpose:** Entry point and HTTP interface (Gin). Connects requests to usecases.
-   **Subfolders:**
    -   `controllers/` — Gin handlers/controllers.
    -   `routers/` — Defines public and protected routes.
    -   `main.go` — Initializes router, DB, services, and starts the server.

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
Fetch and display all tasks currently in the MongoDB store. This endpoint is not protected.

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks`

**Response:**

-   **Status Code:** 200 OK

```json
{
    "Tasks": [
        {
            "id": "64e4f1d9bc19d2e91f6dcaa1",
            "title": "Task 4",
            "description": "Backend with Go",
            "status": "In progress",
            "due_date": "2025-07-21T12:00:00Z"
        }
    ]
}
```

**Errors:**

-   `500 Internal Server Error`:

```json
{ "error": "Database failure" }
```

---

### 2️⃣ GET `/tasks/:id`

**Description:**  
Fetch a single task by its MongoDB ObjectID. Not protected.

**Request Example:**
`http://localhost:8080/tasks/64e4f1d9bc19d2e91f6dcaa1`

**Response:**

-   **Status Code:** 200 OK

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

-   `404 Not Found`
-   `500 Internal Server Error`

---

### 3️⃣ PUT `/tasks/:id`

**Description:**  
Update task (requires authentication).

**Request Body:**

```json
{
    "title": "Task 4 Updated",
    "description": "Updated backend with Go and Gin",
    "status": "Completed",
    "dueDate": "2025-07-21T12:00:00Z"
}
```

**Response:**

```json
{ "message": "Task updated Successfully!" }
```

**Errors:**

-   400, 404, 500

---

### 4️⃣ POST `/tasks`

**Description:**  
Create a new task (requires authentication).

**Request Body:**

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

```json
{
    "_id": "64e4f2b0bc19d2e91f6dca11",
    "title": "Write Documentation",
    "description": "Create API documentation using Markdown",
    "status": "Pending",
    "dueDate": "2025-07-25T23:59:59Z"
}
```

---

### 5️⃣ DELETE `/tasks/:id`

**Description:**  
Delete a task (requires authentication).

**Response:**

```json
{ "message": "Task deleted!" }
```

---

### 6️⃣ POST `/register`

**Description:**  
Register a new user.

**Request Example:**

```json
{
    "id": "1",
    "name": "Mahder Ashenafi",
    "username": "mahderadmin",
    "password": "SecureP@ss123",
    "role": "admin"
}
```

**Response:**

```json
{ "message": "user registered successfully" }
```

---

### 7️⃣ POST `/login`

**Description:**  
Login and receive JWT.

```json
{
    "username": "mahderadmin",
    "password": "SecureP@ss123"
}
```

**Response:**

```json
{
    "message": "user logged in successfully",
    "token": "<JWT_TOKEN>"
}
```

---

### 8️⃣ GET `/user_profile`

**Description:**  
Access only for authenticated users.

```json
{ "message": "Only logged in users can see this" }
```

---

### 9️⃣ GET `/admin_page`

**Description:**  
Access only for users with `admin` role.

```json
{ "message": "Hello, welcome to the admin page!" }
```

---

## 🔐 Middlewares

### `AuthMiddleWare`

-   Verifies JWT from the `Authorization` header.
-   Adds `role` and `username` to context.

### `AuthRoleMiddleWare`

-   Verifies if user has `admin` role from context.

---

## 🧪 Testing

-   Tested using Postman (VS Code extension).
-   All routes return correct status codes and data.

---

## 🚀 Running the API

```bash
go mod tidy
go run main.go
```

---

## 🧩 MongoDB Integration

### .env Example

```
MONGODB_URI=mongodb://localhost:27017
DB_NAME=taskmanager
```

### MongoDB Connection

```go
client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
db := client.Database(os.Getenv("DB_NAME"))
```

---

## 📘 Update Summary

-   Switched from in-memory store to MongoDB.
-   Modularized into Clean Architecture layers.
-   Middleware-based authentication and role protection added.
