# WebKeeper

A REST API for monitoring website availability and tracking HTTP status codes over time.

## Features

- Register websites to monitor
- Automatic health checks every 5 seconds (configurable)
- Logs HTTP status codes with timestamps
- RESTful API with Gin framework
- PostgreSQL persistence with GORM
- Docker Compose for local development
- Concurrent site checking with goroutines

## Prerequisites

- Go 1.22+
- PostgreSQL database (or use Docker Compose)

## Quick Start

### Using Docker Compose

```bash
# Start PostgreSQL and pgAdmin
docker compose up -d

# Run the application
go run cmd/main.go
```

pgAdmin is available at http://localhost:15432 (admin@admin.com / my_password)

### Manual Setup

Set environment variables and run:

```bash
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=webkeeper
go run cmd/main.go
```

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | localhost | Database host |
| `DB_USER` | postgres | Database user |
| `DB_PASSWORD` | my_password | Database password |
| `DB_NAME` | polls | Database name |
| `DB_PORT` | 5432 | Database port |
| `PORT` | 3000 | HTTP server port |
| `MONITOR_INTERVAL` | 5 | Check interval in seconds |

## API Endpoints

### List All Sites

```bash
GET /sites
```

Response:
```json
[
  {"id": 1, "nome": "Google", "url": "https://google.com"},
  {"id": 2, "nome": "GitHub", "url": "https://github.com"}
]
```

### Get Site by ID

```bash
GET /sites/:id
```

### Get Site Logs

```bash
GET /sites/:id/logs
```

Response:
```json
[
  {"id": 1, "site_id": 1, "status_code": 200, "log_des": "OK", "hora": "2024-12-20T10:00:00Z"},
  {"id": 2, "site_id": 1, "status_code": 200, "log_des": "OK", "hora": "2024-12-20T10:00:05Z"}
]
```

### Create Site

```bash
POST /sites
Content-Type: application/json

{
  "nome": "Google",
  "url": "https://google.com"
}
```

### Update Site

```bash
PATCH /sites/:id
Content-Type: application/json

{
  "nome": "New Name",
  "url": "https://new-url.com"
}
```

### Delete Site

```bash
DELETE /sites/:id
```

## Project Structure

```
WebKeeper/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── controllers/
│   │   └── controllers.go      # HTTP request handlers
│   ├── database/
│   │   └── db.go               # Database connection
│   ├── models/
│   │   └── models.go           # Site and LogSite models
│   ├── routes/
│   │   └── routes.go           # API route definitions
│   └── services/
│       ├── checker.go          # Site monitoring service
│       └── logs.go             # Log persistence
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Data Models

### Site

| Field | Type | Description |
|-------|------|-------------|
| ID | uint | Primary key |
| Nome | string | Site name |
| URL | string | URL to monitor (unique) |

### LogSite

| Field | Type | Description |
|-------|------|-------------|
| ID | uint | Primary key |
| SiteID | uint | Foreign key to Site |
| StatusCode | int | HTTP status code |
| LogDes | string | Status description |
| Hora | time.Time | Timestamp |

## Monitoring

The application starts a background goroutine that:

1. Fetches all registered sites from the database
2. Makes HTTP GET requests to each site concurrently
3. Records the status code and description
4. Waits for the configured interval
5. Repeats indefinitely

Connection errors are logged with status code 0 and the error message.
