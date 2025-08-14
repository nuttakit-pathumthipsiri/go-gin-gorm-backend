# Go Gin GORM Backend

A RESTful API for managing topics and topic details using Go, Gin, and GORM.

## Prerequisites

- Go 1.19+
- SQL Server (Docker)
- Swag CLI
- Make

### Installing Make

**Windows:**

```bash
# Using Chocolatey
choco install make

```

**macOS:**

```bash
# Using Homebrew
brew install make

```

## Quick Start

### 1. Setup Database

Connect to SQL Server container:

```bash
docker-compose up
```

### 2. Run Application

Generate Swagger docs and start the server:

```bash
make run dev
```

### 3. API Documentation

Access Swagger UI at: http://localhost:8080/swagger/index.html

## Available Make Commands

- `make run dev` - Generate Swagger docs and run the application
- `make run` - Run the application only
- `make swagger` - Generate Swagger documentation
- `make build` - Build the application

## Environment Variables

Copy the `.env.example` file to `.env` in the root directory and update it with your database configuration.
