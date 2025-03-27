# Task Management API Documentation

This documentation outlines all endpoints of the Task Management API along with sample requests and responses.  
**Note:** This API uses MongoDB as its data store. Ensure that your MongoDB instance is running (default: `mongodb://localhost:27017`) and that the database and collection (`taskdb` and `tasks`) exist or will be created on first use. Task IDs are now represented as MongoDB ObjectIDs – 24-character hexadecimal strings.

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
        "id": "60a6b2fb1234567890abcdef",
        "title": "Example Task 1",
        "description": "This is a sample task description.",
        "due_date": "2025-03-30T00:00:00Z",
        "status": "pending"
    },
    {
        "id": "60a6b2fbabcdef1234567890",
        "title": "Example Task 2",
        "description": "Another task description.",
        "due_date": "2025-04-05T00:00:00Z",
        "status": "completed"
    }
]
```

---

## 2. GET /tasks/:id

**Description:**  
Retrieves details of a specific task by its ID. Remember to use a valid MongoDB ObjectID (24-character hex string).

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks/60a6b2fb1234567890abcdef`  
    Replace with the actual task ID.

**Response:**

-   **Status Code:** 200 OK (on success)
-   **Body Example:**

```json
{
    "id": "60a6b2fb1234567890abcdef",
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
Creates a new task. The provided task data will be stored in MongoDB and an ObjectID will be generated.

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
-   **Body Example:** (The response JSON includes a generated MongoDB ObjectID.)

```json
{
    "id": "60a6b2fbabcdef1234567890",
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
Updates an existing task. Use a valid MongoDB ObjectID for the task ID.

**Request:**

-   **Method:** PUT
-   **URL:** `http://localhost:8080/tasks/60a6b2fbabcdef1234567890`  
    Replace with the actual task ID to update.
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
    "id": "60a6b2fbabcdef1234567890",
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
Deletes a task by its ID. Provide a valid MongoDB ObjectID.

**Request:**

-   **Method:** DELETE
-   **URL:** `http://localhost:8080/tasks/60a6b2fbabcdef1234567890`  
    Replace with the actual task ID.

**Response:**

-   **Status Code:** 204 No Content
-   **Body:** No response body.

-   **Error Responses:**
    -   **400 Bad Request:** Invalid task ID.
    -   **404 Not Found:** Task not found.

---

## Additional MongoDB Integration Details

-   **MongoDB Driver:** This API uses the official MongoDB Go Driver (`go.mongodb.org/mongo-driver/mongo`).
-   **Database & Collection:** The application connects to a MongoDB instance at `mongodb://localhost:27017` by default, uses the database `taskdb` and the collection `tasks`.
-   **ObjectID:** Task IDs are managed by MongoDB as ObjectIDs. When creating tasks, a new ObjectID is generated and included in the response.

---

## Testing with Postman

1. **Create a New Request:**  
   Set the request method (GET, POST, PUT, DELETE) and URL based on the endpoint you are testing.

2. **Set Headers:**  
   For POST and PUT requests, add `Content-Type: application/json`.

3. **Provide Request Body:**  
   For POST and PUT, input the JSON payload as shown in the examples above.

4. **Send the Request and Inspect the Response:**  
   Ensure that the returned values (including generated MongoDB ObjectIDs) match the expected output.

5. **Verify in MongoDB:**  
   Optionally, use MongoDB Compass or the Mongo shell to directly confirm that tasks are being created, updated, and deleted in the database.
