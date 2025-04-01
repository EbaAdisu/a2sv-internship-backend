# Task Manager API

This is a Task Management API backend written in Go using Clean Architecture principles. The API provides endpoints for managing tasks and users with role-based access (admin and regular user) and uses JWT for authentication.

## Table of Contents

-   [Overview](#overview)
-   [Architecture](#architecture)
-   [Folder Structure](#folder-structure)
-   [Setup & Running](#setup--running)
-   [API Endpoints](#api-endpoints)
-   [Contributing](#contributing)
-   [License](#license)

## Overview

The Task Manager API is built according to Clean Architecture principles. The application is divided into distinct layers—Delivery, Domain, Infrastructure, Repositories, and Usecases—to enforce separation of concerns, improve maintainability, testability, and scalability.

## Architecture

-   **Domain:**  
    Contains the core business entities (Task, User) that are independent of external frameworks or libraries.

-   **Usecases:**  
    Implements the application-specific business rules. This layer orchestrates the interactions between domain entities and external dependencies (data persistence, password hashing, JWT generation) by interacting through interfaces.

-   **Repositories:**  
    Provides the data access logic. Interfaces and implementations for MongoDB operations are defined here to abstract database interactions.

-   **Infrastructure:**  
    Implements external dependencies such as JWT token management, password hashing, and HTTP middleware for authentication and authorization.

-   **Delivery:**  
    Contains HTTP handlers (controllers) and routing configuration using the Gin framework. This layer is responsible for receiving and responding to HTTP requests.

## Folder Structure

```

task-manager/
├── Delivery/
│ ├── main.go # Application entry point: initializes dependencies, router, and starts the server.
│ ├── controllers/
│ │ └── controller.go # HTTP handlers that process requests, invoke use cases, and return responses.
│ └── routers/
│ └── router.go # Defines API endpoints and applies necessary middleware.
├── Domain/
│ └── domain.go # Domain models representing the core business entities: Task and User.
├── Infrastructure/
│ ├── auth_middleware.go # JWT-based authentication and authorization middleware.
│ ├── jwt_service.go # Functions for generating and validating JWT tokens.
│ └── password_service.go # Password hashing and verification using bcrypt.
├── Repositories/
│ ├── task_repository.go # Data access layer for tasks.
│ └── user_repository.go # Data access layer for users.
└── Usecases/
├── task_usecases.go # Business logic for tasks, including CRUD operations and get task by ID.
└── user_usecases.go # Business logic for user registration and login.

```

## Setup & Running

### Prerequisites

-   Go 1.16 or later
-   MongoDB running locally or a valid MongoDB URI

### Steps

1. **Clone the Repository:**

    ```shell
    git clone https://github.com/yourusername/task-manager.git
    cd task-manager
    ```

````

2. **Install Dependencies:**

    Ensure Go modules are enabled. Then run:

    ```shell
    go mod tidy
    ```

3. **Run the Server:**

    ```shell
    go run Delivery/main.go
    ```

    The server will start on port 8080 by default.

## API Endpoints

### Public Endpoints

-   **POST /register**
    Registers a new user.

    **Request Body Example:**

    ```json
    {
        "username": "john_doe",
        "password": "securepassword"
    }
    ```

-   **POST /login**
    Logs in a user and returns a JWT token.

    **Request Body Example:**

    ```json
    {
        "username": "john_doe",
        "password": "securepassword"
    }
    ```

### Protected Endpoints (Requires JWT Token)

-   **GET /tasks**
    Retrieves all tasks.

-   **GET /tasks/:id**
    Retrieves a single task by its ID.

### Admin-Only Endpoints

-   **POST /tasks**
    Creates a new task.

    **Request Body Example:**

    ```json
    {
        "title": "Finish Documentation",
        "description": "Complete the API documentation and update the README.",
        "due_date": "2025-04-02T15:04:05Z",
        "status": "pending"
    }
    ```

-   **PUT /tasks/:id**
    Updates an existing task.

-   **DELETE /tasks/:id**
    Deletes a task.

> **Note:** The first registered user is automatically assigned the "admin" role, while subsequent users receive the "user" role by default.

## Contributing

Contributions are welcome.

1. Fork the repository.
2. Create your feature branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin feature/your-feature`
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

```
````
