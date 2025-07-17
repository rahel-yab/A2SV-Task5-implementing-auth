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

## Endpoints

### 1. Get All Tasks

- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieve a list of all tasks.
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

### 2. Get Task by ID

- **URL:** `/tasks/{id}`
- **Method:** `GET`
- **Description:** Retrieve a single task by its ID.
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

### 3. Create a New Task

- **URL:** `/tasks`
- **Method:** `POST`
- **Description:** Add a new task.
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

---

### 4. Update a Task

- **URL:** `/tasks/{id}`
- **Method:** `PUT`
- **Description:** Update an existing task by ID.
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

- **Response Example (Not Found):**

```json
{
  "message": "task not found"
}
```

---

### 5. Delete a Task

- **URL:** `/tasks/{id}`
- **Method:** `DELETE`
- **Description:** Remove a task by its ID.
- **Response Example (Success):**

```json
{
  "message": "Task removed"
}
```

- **Response Example (Not Found):**

```json
{
  "message": "task not found"
}
```

---

## Task Object

| Field       | Type   | Description        |
| ----------- | ------ | ------------------ |
| id          | string | Unique identifier  |
| title       | string | Title of the task  |
| description | string | Task details       |
| due_date    | string | Due date           |
| status      | string | Task status        |

---

## Error Responses

- All errors return a JSON object with an `error` or `message` field describing the issue.
