# Task-Manager

Welcome to the Task Manager API documentation. This API allows users to manage tasks efficiently by providing endpoints to create, read, update, and delete tasks. The API is designed to handle task-related operations such as setting task priorities, tracking due dates, and updating task statuses. Built with Go and utilizing the Gin framework, the Task Manager API ensures high performance and scalability for your task management needs.

##Get All Tasks
`{{URL}}api/v1/tasks`

## API Request Description

This API endpoint makes an HTTP GET request to retrieve a list of tasks.

### Request Parameters

- No request parameters are required for this endpoint.

### Query Parameters

- No query parameters are required for this endpoint.

### Authentication

- The user needs to be authenticated to access this endpoint.

## API Response

The API returns a JSON object containing a list of tasks. Each task is represented by an object with the following properties:

- `id` (number): The unique identifier of the task.
- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `status` (string): The status of the task.
- `priority` (string): The priority of the task.
- `due_date` (string): The due date of the task.
- `created_at` (string): The timestamp when the task was created.
- `updated_at` (string): The timestamp when the task was last updated.

### Response Example

```json
{
  "2": {
    "id": 0,
    "title": "",
    "description": "",
    "status": "",
    "priority": "",
    "due_date": "",
    "created_at": "",
    "updated_at": ""
  },
  "3": {
    "id": 0,
    "title": "",
    "description": "",
    "status": "",
    "priority": "",
    "due_date": "",
    "created_at": "",
    "updated_at": ""
  }
}
```

## Create Task

`{{URL}}api/v1/tasks/`

### Create Task

This endpoint allows the user to create a new task.

#### Request Body

- `id` (number, required): The unique identifier for the task.
- `title` (string, required): The title of the task.
- `description` (string, required): A brief description of the task.
- `status` (string, required): The status of the task (e.g., completed, pending, in progress).
- `priority` (string, required): The priority level of the task (e.g., high, medium, low).
- `due_date` (string, required): The due date for the task in the format "YYYY-MM-DDTHH:MM:SSZ".
- `created_at` (string, required): The date and time when the task was created in the format "YYYY-MM-DDTHH:MM:SSZ".
- `updated_at` (string, required): The date and time when the task was last updated in the format "YYYY-MM-DDTHH:MM:SSZ".

#### Response (JSON Schema)

```json
{
  "type": "object",
  "properties": {
    "id": { "type": "number" },
    "title": { "type": "string" },
    "description": { "type": "string" },
    "status": { "type": "string" },
    "priority": { "type": "string" },
    "due_date": { "type": "string" },
    "created_at": { "type": "string" },
    "updated_at": { "type": "string" }
  }
}
```

### Body (raw json)

```json
{
  "id": 432,
  "title": "Complete project report",
  "description": "Finish the annual project report by the end of the week",
  "status": "completed",
  "priority": "high",
  "due_date": "0001-01-01T00:00:00Z",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```

##Get Single Task
`{{URL}}api/v1/tasks/:id`

### Request

- Method: GET
- Endpoint: {{URL}}api/v1/tasks/:id

### Response

- Status: 200 OK
- Content-Type: application/json
- Body:
  ```json
  {
    "id": 0,
    "title": "",
    "description": "",
    "status": "",
    "priority": "",
    "due_date": "",
    "created_at": "",
    "updated_at": ""
  }
  ```

This endpoint returns the details of a task identified by the provided ID, including its title, description, status, priority, due date, creation timestamp, and last update timestamp.

#### JSON Schema

```json
{
  "type": "object",
  "properties": {
    "id": {
      "type": "number"
    },
    "title": {
      "type": "string"
    },
    "description": {
      "type": "string"
    },
    "status": {
      "type": "string"
    },
    "priority": {
      "type": "string"
    },
    "due_date": {
      "type": "string"
    },
    "created_at": {
      "type": "string"
    },
    "updated_at": {
      "type": "string"
    }
  }
}
```

### Path Variables

`id              {{randomTen}}`

## Update Single Task

`{{URL}}api/v1/tasks/:id`

### Update Task

This endpoint is used to update a specific task by its ID.

#### Request

- Method: PUT
- URL: `{{URL}}api/v1/tasks/:id`
- Body (raw, application/json):
  ```json
  {
    "id": 0,
    "title": "",
    "description": "",
    "status": "",
    "priority": "",
    "due_date": "",
    "created_at": "",
    "updated_at": ""
  }
  ```

#### Response

The response is a JSON object with the following schema:

```json
{
  "type": "object",
  "properties": {
    "id": { "type": "number" },
    "title": { "type": "string" },
    "description": { "type": "string" },
    "status": { "type": "string" },
    "priority": { "type": "string" },
    "due_date": { "type": "string" },
    "created_at": { "type": "string" },
    "updated_at": { "type": "string" }
  }
}
```

### Path Variables

`id                 {{randomTen}}`

### Body

##### raw (json)

```json
{
  "id": 2,
  "title": "Write unit tests",
  "description": "Write unit tests for the new features implemented",
  "status": "pending",
  "priority": "",
  "due_date": "0001-01-01T00:00:00Z",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```

## Delete Single Task

`{{URL}}api/v1/tasks/:id`

### Delete Task

This endpoint is used to delete a specific task by providing the task ID in the URL.

#### Request

- Method: DELETE
- Headers: No specific headers required
- URL: {{URL}}api/v1/tasks/:id

#### Response

- Status: 200
- Content-Type: application/json
- Body:

  ```json
  {
    "Success": ""
  }
  ```

### Path Variables

`id                 {{randomTen}}`
