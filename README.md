# GraphQL API in Go

This project demonstrates a simple GraphQL API implemented in the Go programming language. The API allows users to query a list of tasks, each having an ID, title, and a completion status.

---

## Features

- **GraphQL Schema**: Defines a `Task` type and a query for fetching tasks.
- **Endpoints**: A single endpoint `/graphql` to handle GraphQL queries.
- **In-Memory Data**: Tasks are stored in memory for simplicity.

---

## Prerequisites

Ensure the following are installed on your system:

- [Go](https://go.dev/dl/) (version 1.20+ recommended)
- A GraphQL client or tool (Postman, Curl, or GraphiQL)

---

## Installation and Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/graphql-go.git
   cd graphql-go

2. **Initialize Dependencies**:
    ```bash
    go mod tidy

3. **Run the Server**:
    ```bash
    go run main.go

- The server will start at http://localhost:8080/graphql.

## Usage

    GraphQL Query
    To interact with the API, you can use Postman, Curl, or a GraphQL client.

- Endpoint: http://localhost:8080/graphql

- Headers:
    Content-Type: application/json
    - Example Query
        ```bash
        query {
        tasks {
            id
            title
            done
        }
        }

    - Expected Response
        ```bash
        {
        "data": {
            "tasks": [
            {
                "id": "1",
                "title": "Learn GraphQL",
                "done": false
            },
            {
                "id": "2",
                "title": "Build a GraphQL API in Go",
                "done": false
            }
            ]
        }
        }


## Project Structure

        graphql-go/
        │
        ├── main.go          # Main server file
        ├── go.mod           # Go module dependencies
        └── go.sum           # Go module checksums
