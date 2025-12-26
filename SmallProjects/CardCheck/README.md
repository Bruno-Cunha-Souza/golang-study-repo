# CardCheck

A simple HTTP API for validating credit card numbers using the Luhn algorithm.

## Features

- Validates credit card numbers using the Luhn checksum algorithm
- RESTful HTTP API
- Input validation (digits only, minimum length)
- Configurable port

## Usage

### Running the Server

```bash
# Default port (8080)
go run main.go

# Custom port
go run main.go 3000
```

### API Endpoint

**POST /**

Validates a credit card number.

**Request:**
```json
{
  "number": "4532015112830366"
}
```

**Response (valid):**
```json
{
  "valid": true
}
```

**Response (invalid):**
```json
{
  "valid": false
}
```

**Error Response (bad input):**
```
400 Bad Request
invalid input: must contain only digits and be at least 2 characters
```

### Examples

```bash
# Valid card number
curl -X POST -d '{"number":"4532015112830366"}' http://localhost:8080

# Invalid card number
curl -X POST -d '{"number":"1234567890"}' http://localhost:8080

# Invalid input (contains non-digits)
curl -X POST -d '{"number":"4111-1111-1111-1111"}' http://localhost:8080
```

## Project Structure

```
CardCheck/
├── main.go          # HTTP server and request handler
├── luhn/
│   └── luhn.go      # Luhn algorithm implementation
├── go.mod
└── README.md
```

## How the Luhn Algorithm Works

1. Starting from the rightmost digit, double every second digit
2. If doubling results in a number > 9, subtract 9
3. Sum all digits
4. If the total modulo 10 equals 0, the number is valid
