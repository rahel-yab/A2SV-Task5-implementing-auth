# Task Management REST API

A simple RESTful API for managing tasks, now with persistent storage using MongoDB.

## Features

- Create, read, update, and delete tasks
- Data is persisted in MongoDB (local or cloud)
- Built with Go and Gin

## Prerequisites

- Go 1.18+
- MongoDB (local installation or MongoDB Atlas cloud instance)

## Setup

### 1. Clone the Repository

```sh
git clone <your-repo-url>
cd <repo-directory>
```

### 2. Install Dependencies

```sh
go mod tidy
```

### 3. Configure MongoDB Connection

Set the `MONGODB_URI` environment variable to your MongoDB connection string.

- **Local MongoDB (default):**
  - The app will use `mongodb://localhost:27017` if `MONGODB_URI` is not set.
- **Custom/Cloud MongoDB:**
  - **Linux/macOS:**
    ```sh
    export MONGODB_URI="your-mongodb-uri"
    ```
  - **Windows PowerShell:**
    ```powershell
    $env:MONGODB_URI="your-mongodb-uri"
    ```

### 4. Run the Application

```sh
go run task_manager/main.go
```

The API will be available at `http://localhost:8080` by default.

## API Endpoints

See [`docs/api_documentation.md`](./docs/api_documentation.md) for full endpoint documentation and examples.

## Verifying Data in MongoDB

- Use the MongoDB shell (`mongosh`) or MongoDB Compass to inspect your data:
  1. Connect: `mongosh`
  2. Switch DB: `use taskdb`
  3. Show tasks: `db.tasks.find().pretty()`

## Example Task JSON

```json
{
  "id": "1",
  "title": "Sample Task",
  "description": "This is a sample task.",
  "due_date": "2024-07-18T00:00:00Z",
  "status": "pending"
}
```

