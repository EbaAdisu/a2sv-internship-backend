# Task Management API Documentation

This documentation outlines all endpoints of the Task Management API along with sample requests and responses.  
**Note:**

-   This API uses JWT-based authentication and MongoDB as its datastore.
-   Ensure your MongoDB instance is running at `mongodb://localhost:27017`.
-   The application uses the database `taskdb` and the collections `tasks` and `users`.
-   Task IDs and User IDs are represented as MongoDB ObjectIDs – 24-character hexadecimal strings.

---

## User Authentication

### Register

**Description:**  
Creates a new user account. The very first user created will be assigned the role of `admin`; subsequent users will have the role `user` by default.

**Request:**

-   **Method:** POST
-   **URL:** `http://localhost:8080/register`
-   **Headers:** `Content-Type: application/json`
-   **Request Body Example:**

```json
{
    "username": "johndoe",
    "password": "securepassword"
}
```

**Response:**

-   **Status Code:** 201 Created
-   **Body Example:**

```json
{
    "id": "60a6b2fbabcdef1234567890",
    "username": "johndoe",
    "role": "admin"
}
```

### Login

**Description:**  
Authenticates a user and returns a JWT token upon successful login.

**Request:**

-   **Method:** POST
-   **URL:** `http://localhost:8080/login`
-   **Headers:** `Content-Type: application/json`
-   **Request Body Example:**

```json
{
    "username": "johndoe",
    "password": "securepassword"
}
```

**Response:**

-   **Status Code:** 200 OK
-   **Body Example:**

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

## Task Endpoints

**Authentication Requirement:**  
All task endpoints require a valid JWT token provided in the `Authorization` header in the following format:  
er`Authorization: Bear <token>`

### Get All Tasks

**Description:**  
Retrieves a list of all tasks. Accessible by any authenticated user.

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks`
-   **Headers:**
    -   `Authorization: Bearer <token>`

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

### Get Task by ID

**Description:**  
Retrieves details of a specific task by its ID. Use a valid MongoDB ObjectID.

**Request:**

-   **Method:** GET
-   **URL:** `http://localhost:8080/tasks/60a6b2fb1234567890abcdef`  
    _(Replace with the actual task ID)_
-   **Headers:**
    -   `Authorization: Bearer <token>`

**Response:**

-   **Status Code:** 200 OK
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

### Create Task (Admin Only)

**Description:**  
Creates a new task. Only authenticated users with the `admin` role can access this endpoint.

**Request:**

-   **Method:** POST
-   **URL:** `http://localhost:8080/tasks`
-   **Headers:**
    -   `Authorization: Bearer <token>`
    -   `Content-Type: application/json`
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
-   **Body Example:**

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
    -   **403 Forbidden:** Insufficient privileges (if non-admin).

### Update Task (Admin Only)

**Description:**  
Updates an existing task. Only accessible by users with the `admin` role.

**Request:**

-   **Method:** PUT
-   **URL:** `http://localhost:8080/tasks/60a6b2fbabcdef1234567890`  
    _(Replace with the actual task ID)_
-   **Headers:**
    -   `Authorization: Bearer <token>`
    -   `Content-Type: application/json`
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
-   **Body Example:**

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
    -   **403 Forbidden:** Insufficient privileges (if non-admin).

### Delete Task (Admin Only)

**Description:**  
Deletes a task by its ID. Only accessible by users with the `admin` role.

**Request:**

-   **Method:** DELETE
-   **URL:** `http://localhost:8080/tasks/60a6b2fbabcdef1234567890`  
    _(Replace with the actual task ID)_
-   **Headers:**
    -   `Authorization: Bearer <token>`

**Response:**

-   **Status Code:** 204 No Content
-   **Body:** No response body.

-   **Error Responses:**
    -   **400 Bad Request:** Invalid task ID.
    -   **404 Not Found:** Task not found.
    -   **403 Forbidden:** Insufficient privileges (if non-admin).

---

## Additional MongoDB Integration Details

-   **MongoDB Driver:**  
    This API uses the official MongoDB Go Driver (`go.mongodb.org/mongo-driver/mongo`).
-   **Database & Collection:**  
    The application connects to a MongoDB instance at `mongodb://localhost:27017`, uses the database `taskdb` and the collections `tasks` and `users`.
-   **ObjectID:**  
    Task and User IDs are managed by MongoDB as ObjectIDs. When creating tasks or users, a new ObjectID is generated and included in the response.

---

## Testing with Postman

1. **User Registration and Login:**
    - Create a new request with method POST to `http://localhost:8080/register` to create a new account.
    - Use POST to `http://localhost:8080/login` with your credentials to receive a JWT token.
2. **Using JWT Token:**
    - For protected endpoints (tasks), add the header:  
      `Authorization: Bearer <token>` (with your received token).
3. **Task Management:**
    - Test GET, POST, PUT, DELETE endpoints as described above.
    - Verify responses and match ObjectID formats.
4. **Check MongoDB Directly:**
    - Optionally, use MongoDB Compass or the Mongo shell to verify that your tasks and users are stored correctly.
