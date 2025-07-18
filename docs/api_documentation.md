# Task Management REST API Documentation

## MongoDB Integration

This API now uses MongoDB for persistent data storage. You must have a running MongoDB instance (local or cloud) and configure the connection string.

### MongoDB Setup

- **Local:** Download and install MongoDB Community Edition from https://www.mongodb.com/try/download/community
- **Cloud:** Create a free cluster at https://www.mongodb.com/cloud/atlas and get your connection string.

### Configuration

- Set the environment variable `MONGODB_URI` before running the API:
  - **Linux/macOS:**
    ```sh
    export MONGODB_URI="mongodb://localhost:27017"
    ```
  - **Windows PowerShell:**
    ```powershell
    $env:MONGODB_URI="mongodb://localhost:27017"
    ```
  - If not set, defaults to `mongodb://localhost:27017`.
- The API uses the database `taskdb` and the collection `tasks` by default.

### Verifying Data

- Use the MongoDB shell (`mongosh`) or MongoDB Compass to inspect your data:
  1. Connect: `mongosh`
  2. Switch DB: `use taskdb`
  3. Show tasks: `db.tasks.find().pretty()`

---

## Base URL

```
http://localhost:8080
```

---

## Authentication & Authorization

### JWT Authentication

- Most endpoints require a valid JWT token in the `Authorization` header:
  ```
  Authorization: Bearer <your_jwt_token>
  ```
- Obtain a token via the `/login` endpoint.
- The token contains user information and role ("admin" or "user").

### User Roles

- **admin**: Can create, update, delete, and promote users. Can view all tasks.
- **user**: Can view all tasks and get task by ID.
- The first registered user is automatically assigned the admin role.
- Only admins can promote other users to admin.

---

## Endpoints

### 1. Register a New User

- **URL:** `/register`
- **Method:** `POST`
- **Description:** Create a new user account. The first user becomes admin, others are regular users.
- **Request Body:**

```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "yourpassword"
}
```

- **Response Example:**

```json
{
  "message": "User registered successfully",
  "role": "admin"
}
```

- **Notes:**
  - Usernames and emails must be unique.

---

### 2. Login

- **URL:** `/login`
- **Method:** `POST`
- **Description:** Authenticate a user and receive a JWT token.
- **Request Body:**

```json
{
  "username": "johndoe", // or
  "email": "john@example.com",
  "password": "yourpassword"
}
```

- **Response Example:**

```json
{
  "message": "User logged in successfully",
  "token": "<jwt_token>",
  "role": "admin"
}
```

- **Notes:**
  - You can log in with either username or email (password required).

---

### 3. Promote User to Admin

- **URL:** `/promote`
- **Method:** `POST`
- **Authentication:** Admin only (JWT required)
- **Description:** Promote a user to admin by username or email.
- **Request Header:**
  - `Authorization: Bearer <jwt_token>`
- **Request Body:**

```json
{
  "identifier": "johndoe" // username or email
}
```

- **Response Example:**

```json
{
  "message": "User promoted to admin"
}
```

- **Error Example (not admin):**

```json
{
  "error": "Admin access required"
}
```

---

### 4. Get All Tasks

- **URL:** `/tasks`
- **Method:** `GET`
- **Authentication:** Any authenticated user (JWT required)
- **Description:** Retrieve a list of all tasks.
- **Request Header:**
  - `Authorization: Bearer <jwt_token>`
- **Response Example:**

```json
[
  {
    "id": "1",
    "title": "Sample Task",
    "description": "This is a sample",
    "due_date": "2024-06-01T00:00:00Z",
    "status": "pending"
  }
]
```

---

### 5. Get Task by ID

- **URL:** `/tasks/{id}`
- **Method:** `GET`
- **Authentication:** Any authenticated user (JWT required)
- **Description:** Retrieve a single task by its ID.
- **Request Header:**
  - `Authorization: Bearer <jwt_token>`
- **Response Example (Success):**

```json
{
  "id": "1",
  "title": "Sample Task",
  "description": "This is a sample",
  "due_date": "2024-06-01T00:00:00Z",
  "status": "pending"
}
```

- **Response Example (Not Found):**

```json
{
  "message": "task not found"
}
```

---

### 6. Create a New Task

- **URL:** `/tasks`
- **Method:** `POST`
- **Authentication:** Admin only (JWT required)
- **Description:** Add a new task.
- **Request Header:**
  - `Authorization: Bearer <jwt_token>`
- **Request Body Example:**

```json
{
  "id": "2",
  "title": "New Task",
  "description": "Details about the new task",
  "due_date": "2024-06-10T00:00:00Z",
  "status": "pending"
}
```

- **Response Example:**

```json
{
  "message": "Task created"
}
```

- **Error Example (not admin):**

```json
{
  "error": "Admin access required"
}
```

---

### 7. Update a Task

- **URL:** `/tasks/{id}`
- **Method:** `PUT`
- **Authentication:** Admin only (JWT required)
- **Description:** Update an existing task by ID.
- **Request Header:**
  - `Authorization: Bearer <jwt_token>`
- **Request Body Example:**

```json
{
  "id": "2",
  "title": "Updated Task",
  "description": "Updated details",
  "due_date": "2024-06-15T00:00:00Z",
  "status": "completed"
}
```

- **Response Example (Success):**

```json
{
  "message": "Task updated"
}
```

- **Error Example (not admin):**

```json
{
  "error": "Admin access required"
}
```

---

### 8. Delete a Task

- **URL:** `/tasks/{id}`
- **Method:** `DELETE`
- **Authentication:** Admin only (JWT required)
- **Description:** Remove a task by its ID.
- **Request Header:**
  - `Authorization: Bearer <jwt_token>`
- **Response Example (Success):**

```json
{
  "message": "Task removed"
}
```

- **Error Example (not admin):**

```json
{
  "error": "Admin access required"
}
```

---

## Task Object

| Field       | Type   | Description       |
| ----------- | ------ | ----------------- |
| id          | string | Unique identifier |
| title       | string | Title of the task |
| description | string | Task details      |
| due_date    | string | Due date          |
| status      | string | Task status       |

---

## User Object

| Field    | Type   | Description          |
| -------- | ------ | -------------------- |
| id       | string | Unique identifier    |
| username | string | Unique username      |
| email    | string | Unique email address |
| password | string | Hashed password      |
| role     | string | "admin" or "user"    |

**Example:**

```json
{
  "id": "60c72b2f9b1e8b001c8e4b8a",
  "username": "johndoe",
  "email": "john@example.com",
  "role": "admin"
}
```

---

## Error Responses

- All errors return a JSON object with an `error` or `message` field describing the issue.

---

## Usage Notes

- Always include the JWT token in the `Authorization` header for protected endpoints.
- Only admins can create, update, delete, or promote users.
- The first user to register is automatically an admin.
- Use `/promote` to grant admin rights to other users (admin only).
- Passwords are securely hashed using bcrypt.
