# Echo Practice (study_go_web)

> **Disclaimer:** This project is for **learning and practice purposes only**. It is used to explore the Go language, the Echo web framework, and clean architecture principles.

## 🚀 Project Overview

A simple RESTful API built with Go and the Echo framework, demonstrating basic CRUD operations, dependency injection, and database interaction using GORM.

### Features
- User Management (Create, Read, Delete)
- Clean Architecture (Handler -> Service -> Repository)
- Configuration management using Viper
- Request validation using `go-playground/validator`
- Database ORM with GORM (MySQL)
- Hot-reloading with Air

## 🛠 Tech Stack
- **Language:** Go
- **Web Framework:** [Echo](https://echo.labstack.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** MySQL
- **Config Management:** [Viper](https://github.com/spf13/viper)
- **Validation:** [Validator](https://github.com/go-playground/validator)

## 🏗 Project Structure
```text
.
├── cmd/server/          # Application entry point
├── configs/             # Configuration files (YAML)
├── internal/
│   ├── config/          # Config loading logic
│   ├── handler/         # HTTP handlers & routing
│   ├── model/           # Database models
│   ├── repository/      # Data access layer
│   └── service/         # Business logic layer
└── pkg/                 # Shared utilities (if any)
```

## 🚦 Getting Started

### Prerequisites
- Go 1.25+
- MySQL instance
- [Air](https://github.com/air-verse/air) (optional, for live reload)

### Installation
1. Clone the repository:
   ```bash
   git clone git@github.com:kingoflongevity/study_go_web.git
   cd study_go_web
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Configure your database in `configs/config.yaml`.

### Running the Application
**With Air (Recommended for development):**
```bash
air
```

**Standard Go run:**
```bash
go run cmd/server/main.go
```

## 🛣 API Endpoints
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create a new user
- `DELETE /api/v1/users/:id` - Delete user by ID
