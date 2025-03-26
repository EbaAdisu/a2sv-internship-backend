Below is a sample API documentation in Markdown that clearly describes each endpoint along with example Postman request details, payloads, and responses.

````markdown
# Task Management API Documentation

This documentation outlines all endpoints of the Task Management API along with sample requests and responses. Use Postman (or a similar tool) to test each endpoint.

---

## 1. GET /tasks

**Description:**  
Retrieves a list of all tasks.

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks`

**Response:**

-   **Status Code:** 200 OK
-   **Body Example:**

```json
[
    {
        "id": 1,
        "title": "Example Task 1",
        "description": "This is a sample task description.",
        "due_date": "2025-03-30T00:00:00Z",
        "status": "pending"
    },
    {
        "id": 2,
        "title": "Example Task 2",
        "description": "Another task description.",
        "due_date": "2025-04-05T00:00:00Z",
        "status": "completed"
    }
]
```
````

---

## 2. GET /tasks/:id

**Description:**  
Retrieves details of a specific task by its ID.

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks/1`  
    Replace `1` with the actual task ID.

**Response:**

-   **Status Code:** 200 OK (on success)
-   **Body Example:**

```json
{
    "id": 1,
    "title": "Example Task 1",
    "description": "This is a sample task description.",
    "due_date": "2025-03-30T00:00:00Z",
    "status": "pending"
}
```

-   **Error Responses:**
    -   **400 Bad Request:** Invalid task ID.
    -   **404 Not Found:** Task not found.

---

## 3. POST /tasks

**Description:**  
Creates a new task.

**Request:**

-   **Method:** POST
-   **URL:** `http://localhost:8080/tasks`
-   **Headers:** `Content-Type: application/json`
-   **Request Body Example:**

```json
{
    "title": "New Task",
    "description": "Details of the new task.",
    "due_date": "2025-04-15T00:00:00Z",
    "status": "pending"
}
```

**Response:**

-   **Status Code:** 201 Created
-   **Body Example:** (The response JSON includes a generated task id.)

```json
{
    "id": 3,
    "title": "New Task",
    "description": "Details of the new task.",
    "due_date": "2025-04-15T00:00:00Z",
    "status": "pending"
}
```

-   **Error Response:**
    -   **400 Bad Request:** Invalid input data.

---

## 4. PUT /tasks/:id

**Description:**  
Updates an existing task.

**Request:**

-   **Method:** PUT
-   **URL:** `http://localhost:8080/tasks/1`  
    Replace `1` with the actual task ID to update.
-   **Headers:** `Content-Type: application/json`
-   **Request Body Example:**

```json
{
    "title": "Updated Task Title",
    "description": "Updated description of the task.",
    "due_date": "2025-04-20T00:00:00Z",
    "status": "completed"
}
```

**Response:**

-   **Status Code:** 200 OK
-   **Body Example:** (Returns the updated task.)

```json
{
    "id": 1,
    "title": "Updated Task Title",
    "description": "Updated description of the task.",
    "due_date": "2025-04-20T00:00:00Z",
    "status": "completed"
}
```

-   **Error Responses:**
    -   **400 Bad Request:** Invalid input or task ID.
    -   **404 Not Found:** Task not found.

---

## 5. DELETE /tasks/:id

**Description:**  
Deletes a task by its ID.

**Request:**

-   **Method:** DELETE
-   **URL:** `http://localhost:8080/tasks/1`  
    Replace `1` with the actual task ID.

**Response:**

-   **Status Code:** 204 No Content
-   **Body:** No response body.

-   **Error Responses:**
    -   **400 Bad Request:** Invalid task ID.
    -   **404 Not Found:** Task not found.

---

## Testing with Postman

1. **Create a new request in Postman.**
2. **Set the request method and URL** according to the endpoint you’re testing.
3. **Add required headers** (e.g., `Content-Type: application/json` for POST and PUT requests).
4. **Provide JSON request body** as shown in the examples for POST and PUT.
5. **Send the request** and inspect the response to ensure it matches the expected output.
