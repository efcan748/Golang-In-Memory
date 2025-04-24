# Golang In-Memory Store

A high-performance in-memory key-value and list store written in Go. Built with a clean RESTful interface, TTL support, and Dockerized for easy deployment.

---

## Features

- Key-Value store with TTL (time-to-live)
- List operations: LPush, LPop
- Fast and lightweight
- REST API
- Docker support
- Integration testable

---

## 📡 REST API Specification

### Installation

#### From Source
 
 Build and run:

```bash
    go run cmd/server/main.go
```

Command Line Options
``` bash
go run cmd/server/main.go [flags]
```

#### Available flags:

- `--host` - Server host address (default: "127.0.0.1")
- `--port` - Server port number (default: "8080")
- `--cleanup` - Cleanup interval in minutes (default: 1)

Example:

```bash
go run cmd/server/main.go --host 0.0.0.0 --port 5000 --cleanup 5
```
### Using Docker

```bash
    docker build -t simple-store .
    docker run -p 8080:8080 simple-store
```


---

### 🔹 Key-Value Endpoints

#### POST `/string`
Set a key-value to store string with optional TTL (milliseconds).

**Request:**
```json
{
  "key": "username",
  "value": "john_doe",
  "ttl": 60000
}
```
Response:
```json
{ "status": "successfully stored." }
```

#### GET /string?key=username
Get value by key.

Response:
```json
{ "value": "john_doe" }
```

#### PUT /string/username
Update a key’s value or TTL.

Request:
```json
{
  "value": "jane_doe",
  "ttl": 120000
}
```
Response

```json
{ "status": "successfully updated." }
```

#### DELETE /string/username
Delete a key.

Response:
```json
{ "status": "successfully removed." }
```

### List Endpoints `/list`

#### POST /list
Push values to a list with TTL.

Request:
```json
{
  "key": "tasks",
  "value": ["task1", "task2"],
  "ttl": 30000
}
```

Response:
```json
{ "status": "successfully stored." }
```

#### GET /list?key=tasks
Pop the last item from the list.

Response:
```json
{ "value": "task1" }
```