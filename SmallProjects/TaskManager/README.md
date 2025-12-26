# TaskManager

A command-line task management application with PostgreSQL persistence.

## Features

- Add tasks with title, description, and deadline
- List all tasks with status and deadline
- Update task status (pending, in progress, completed)
- Delete tasks
- PostgreSQL database with GORM ORM
- Environment variable configuration

## Prerequisites

- Go 1.23+
- PostgreSQL database

## Configuration

Set the following environment variables (or use defaults):

| Variable      | Default           | Description       |
| ------------- | ----------------- | ----------------- |
| `DB_HOST`     | localhost         | Database host     |
| `DB_USER`     | root              | Database user     |
| `DB_PASSWORD` | root              | Database password |
| `DB_NAME`     | root              | Database name     |
| `DB_PORT`     | 5432              | Database port     |
| `DB_TIMEZONE` | America/Sao_Paulo | Timezone          |

## Usage

### Show Help

```bash
go run cmd/main.go help
```

### Add a Task

```bash
go run cmd/main.go add "Task Title" "Task Description" "2024-12-31"
```

The deadline must be in `YYYY-MM-DD` format.

### List Tasks

```bash
go run cmd/main.go list
```

Output:

```
1: Buy groceries - Weekly shopping (Status: pending, Deadline: 2024-12-20)
2: Finish report - Q4 financial report (Status: in_progress, Deadline: 2024-12-25)
```

### Update Task Status

```bash
go run cmd/main.go update [id] [status]
```

Valid statuses:

- `pendente` - Pending
- `em_progresso` - In Progress
- `concluida` - Completed

Example:

```bash
go run cmd/main.go update 1 em_progresso
```

### Delete a Task

```bash
go run cmd/main.go delete [id]
```

## Project Structure

```
TaskManager/
├── cmd/
│   └── main.go       # CLI entry point and command parsing
├── db/
│   └── db.go         # Database connection and configuration
├── handlers/
│   └── task.go       # Business logic (CRUD operations)
├── models/
│   └── task.go       # Task model definition (GORM)
├── go.mod
├── go.sum
└── README.md
```

## Task Model

| Field       | Type      | Description                      |
| ----------- | --------- | -------------------------------- |
| ID          | uint      | Primary key                      |
| Title       | string    | Task title (required)            |
| Description | string    | Task description                 |
| Status      | string    | pending, em_progresso, concluida |
| Deadline    | time.Time | Task deadline                    |
| CreatedAt   | time.Time | Creation timestamp               |
| UpdatedAt   | time.Time | Last update timestamp            |
