# my-task-manager/my-task-manager/README.md

# Task Management REST API

This project is a Task Management REST API built using the Gin Framework in Go. It provides a simple interface for managing tasks with CRUD operations.

## Project Structure

```
my-task-manager
├── controllers
│   └── task_controller.go
├── data
│   └── task_service.go
├── models
│   └── task.go
├── docs
│   └── api_documentation.md
├── main.go
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

-   Go 1.16 or later
-   Gin Framework

### Installation

1. Clone the repository:

    ```
    git clone https://github.com/EbaAdisu/a2sv-internship-backend/
    cd tasks/task5/task-manager
    ```

2. Install dependencies:
    ```
    go mod tidy
    ```

### Running the API

To run the API, execute the following command:

```
go run main.go
```

The server will start on `http://localhost:8080`.

### API Endpoints

-   `GET /tasks` - Retrieve all tasks
-   `GET /tasks/:id` - Retrieve a task by ID
-   `POST /tasks` - Create a new task
-   `PUT /tasks/:id` - Update an existing task
-   `DELETE /tasks/:id` - Delete a task

### Testing

You can use tools like Postman or curl to test the API endpoints.
