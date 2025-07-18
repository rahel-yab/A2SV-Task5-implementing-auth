# Task Management REST API

A simple RESTful API for managing tasks with user authentication and role-based authorization using Go, Gin, and MongoDB.

## Features

- User registration and login (JWT-based)
- Admin and regular user roles
- Only admins can create, update, delete, or promote users
- All authenticated users can view tasks

## Setup

1. Clone the repo and navigate to the project directory.
2. Set up MongoDB and update your connection string if needed.
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

## API Endpoints

- `POST /register` — Register a new user
- `POST /login` — Login and get JWT
- `POST /promote` — Promote user to admin (admin only)
- `GET /tasks` — List all tasks (auth required)
- `POST /tasks` — Create task (admin only)
- `PUT /tasks/:id` — Update task (admin only)
- `DELETE /tasks/:id` — Delete task (admin only)

See `docs/api_documentation.md` for full API details.

