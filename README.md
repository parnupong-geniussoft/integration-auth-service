# integration-auth-service

Authentication service written in Go (Golang) following the Clean Architecture approach. It uses JWT, middleware, and a clear modular separation.

## 🔧 Project Structure

```
.
│   .env                     # Environment variables
│   go.mod / go.sum          # Go module dependencies
│   README.md                # Project documentation
│
├───app
│   └───main.go              # Application entry point
│
├───configs
│   └───config.go            # Load configuration from .env
│
├───docs
│   ├───swagger.yaml/json    # API documentation (Swagger)
│   └───docs.go              # Swagger documentation generation
│
├───modules
│   ├───auth                 # Authentication module
│   │   ├───controllers      # HTTP handlers for auth endpoints
│   │   ├───entities         # Structs / DTOs related to auth
│   │   ├───repositories     # Database interactions
│   │   └───usecases         # Business logic for auth
│   │
│   ├───middlewares          # Middleware
│   └───servers              # Fiber server setup and route handling
│
└───pkg
    ├───databases            # PostgreSQL connection
    ├───loggers              # Logging system
    └───utils                # Helper functions
```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/parnupong-geniussoft/integration-auth-service.git
cd integration-auth-service
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Project

```bash
go run ./app/main.go
```

## 📌 Features

- ✅ JWT Token Authentication

## 🛠 Technologies Used

- [Go Fiber](https://gofiber.io/)
- [JWT (github.com/golang-jwt/jwt)](https://github.com/golang-jwt/jwt)
- [PostgreSQL](https://www.postgresql.org/)
- [Swagger (Swaggo)](https://github.com/swaggo/swag)

## 📄 API Usage

### POST /v1/integration-api/request_token

## 📚 API Documentation

Swagger UI available at:  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## 🧹 Clean Architecture Overview

- `controllers` → Handles requests and invokes usecases
- `usecases` → Business logic implementation
- `repositories` → Handles data persistence
- `entities` → Data structures

## 📁 Middleware

- `middlewares/jwt.go` → JWT Token verification
- `middlewares/loggers.go` → Logs request data into database
- `middlewares/recover.go` → Handles panic and returns safe response
