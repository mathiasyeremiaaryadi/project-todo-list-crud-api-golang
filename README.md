# Golang CRUD Web Blog Post API

This is a CLI application for expense tracker to manage your expenses. 

Project from: https://roadmap.sh/projects/todo-list-api


## Features

- Login
- Register
- Refresh Token
- Create Todo
- Update Todo
- Delete Todo
- Get All Todo
- Get All Todo With Filter
- Get All Todo Sorted DESC


## Documentation

### Login
Endpoint
```
POST /login
```

Request
```
{
  "email": "john@doe.com",
  "password": "password"
}
```

Response
```
HTTP 200/OK
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
}
```

### Register
Endpoint
```
POST /register
```

Request
```
{
  "name": "John Doe",
  "email": "john@doe.com",
  "password": "password"
}
```

Response
```
HTTP 201/Created
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
}
```

### Create Todo
Endpoint
```
POST /todos
```

Request
```
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}
```

Response
```
HTTP 201/Created
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}
```

### Update Todo
Endpoint
```
PUT /todos/1
```

Request
```
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese"
}
```

Response
```
HTTP 200/OK
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese"
}
```

## Delete Todos
Endpoint
```
DELETE /todos/1
```

Request
```
-
```

Response
```
HTTP 204/No Content
-
```

## Get All Todo
Endpoint
```
GET /todos?page=1&limit=10
```

Request
```
-
```

Response
```
{
  "data": [
    {
      "id": 1,
      "title": "Buy groceries",
      "description": "Buy milk, eggs, bread"
    },
    {
      "id": 2,
      "title": "Pay bills",
      "description": "Pay electricity and water bills"
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 2
}
```

## Get All Todo With Filter
Endpoint
```
GET /todos?page=1&limit=1&title=groceries
```

Request
```
-
```

Response
```
{
  "data": [
    {
      "id": 1,
      "title": "Buy groceries",
      "description": "Buy milk, eggs, bread"
    },
  ],
  "page": 1,
  "limit": 10,
  "total": 1
}
```

## Get All Todo Sorted DESC
Endpoint
```
GET /todos?page=1&limit=1&sort=true
```

Request
```
-
```

Response
```
{
  "data": [
    {
      "id": 2,
      "title": "Pay bills",
      "description": "Pay electricity and water bills"
    },
    {
      "id": 1,
      "title": "Buy groceries",
      "description": "Buy milk, eggs, bread"
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 2
}
```


## Clone the project

```bash
git clone https://github.com/mathiasyeremiaaryadi/project-todo-list-crud-api-golang.git
```

Go to the project directory

```bash
cd project-todo-list-crud-api-golang
```

Install dependencies

```bash
go build
```

Start the server

```bash
./todo-list-crud-api
```
