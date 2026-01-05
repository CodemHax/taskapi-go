# RestApI

A simple RESTful API for managing tasks, built with Go, Gin, and Swagger (swaggo).

## Features
- CRUD operations for tasks (Create, Read, Update, Delete)
- In-memory storage (no database required)
- Swagger UI documentation at `/docs`
- Health check endpoint at `/ping`

## Endpoints

### Health Check
- `GET /ping` — Returns `{ "message": "pong" }` if the server is running.

### Tasks
- `GET /tasks` — List all tasks
- `POST /tasks` — Add a new task (JSON body: `{ "id": int, "title": string, "done": bool }`)
- `GET /tasks/{id}` — Get a task by ID
- `PUT /tasks/{id}` — Update a task by ID (JSON body: `{ "title": string, "done": bool }`)
- `DELETE /tasks/{id}` — Delete a task by ID

## Getting Started

### Prerequisites
- Go 1.18 or newer

### Install dependencies
```
go mod tidy
```

### Generate Swagger docs
```
swag init
```

### Run the server
```
go run main.go
```

The API will be available at `http://localhost:8080` and Swagger UI at `http://localhost:8080/docs/index.html`.

## Project Structure
- `main.go` — Main application and route handlers
- `docs/` — Auto-generated Swagger documentation
- `go.mod`, `go.sum` — Go module files

## API Documentation
Visit [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html) after running the server to view and interact with the API documentation.

## License
MIT

