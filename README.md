# Personal Task Manager

[![CI - Unit Tests](https://github.com/VidIsWandering/clean-architecture-golang/actions/workflows/ci.yml/badge.svg)](https://github.com/VidIsWandering/clean-architecture-golang/actions/workflows/ci.yml)
[![Integration Tests](https://github.com/VidIsWandering/clean-architecture-golang/actions/workflows/integration.yml/badge.svg)](https://github.com/VidIsWandering/clean-architecture-golang/actions/workflows/integration.yml)
[![codecov](https://codecov.io/gh/VidIsWandering/clean-architecture-golang/branch/main/graph/badge.svg)](https://codecov.io/gh/VidIsWandering/clean-architecture-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/VidIsWandering/clean-architecture-golang/src)](https://goreportcard.com/report/github.com/VidIsWandering/clean-architecture-golang/src)
[![Go Reference](https://pkg.go.dev/badge/github.com/VidIsWandering/clean-architecture-golang)](https://pkg.go.dev/github.com/VidIsWandering/clean-architecture-golang)
[![Go version](https://img.shields.io/badge/go-1.21-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple Personal Task Manager built with Go using Clean Architecture principles.

## Features

- Create new tasks
- Update task status (TODO → DOING → DONE)
- View tasks by status
- Delete tasks
- REST API interface

## Architecture

This project follows Clean Architecture with the following layers:

- **Domain**: Core business logic and entities
- **Application**: Use cases and application services
- **Infrastructure**: Repository implementations and persistence
- **Presentation**: HTTP controllers and DTOs

## API Endpoints

- `POST /tasks` - Create a new task
- `PUT /tasks/{id}/status` - Update task status
- `GET /tasks?status={status}` - Get tasks by status
- `DELETE /tasks/{id}` - Delete a task

### Example Requests

Create Task:

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "My Task", "description": "Task description"}'
```

Update Status:

```bash
curl -X PUT http://localhost:8080/tasks/123/status \
  -H "Content-Type: application/json" \
  -d '{"newStatus": "doing"}'
```

Get Tasks:

```bash
curl http://localhost:8080/tasks?status=todo
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone the repository:

```bash
git clone https://github.com/VidIsWandering/clean-architecture-golang.git
cd clean-architecture-golang/src
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o task-manager main.go
```

## Project Structure

```
src/
├── domain/
│   ├── entities/
│   └── value_objects/
├── application/
│   ├── usecases/
│   ├── ports/
│   └── dto/
├── infrastructure/
│   ├── persistence/
│   └── repositories/
├── presentation/
│   ├── controllers/
│   └── dto/
├── main.go
└── go.mod
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
