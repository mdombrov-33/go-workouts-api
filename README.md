# Go Fitness Management API

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](https://golang.org/)

A simple REST API for managing fitness workouts and users, built with Go.

## 🚀 Features

- User registration and authentication
- Workout management (CRUD operations)
- PostgreSQL database integration
- Docker Compose for easy database setup

## 🛠 Tech Stack

- **Language**: Go 1.24.5
- **Database**: PostgreSQL (with pgx driver)
- **Web Framework**: Chi v5
- **Migration Tool**: Goose v3
- **Password Hashing**: bcrypt
- **Testing**: testify
- **Containerization**: Docker Compose

## 📦 Installation

1. Ensure you have Go installed on your system.
2. Clone the repository:
   ```bash
   git clone https://github.com/mdombrov-33/go-workouts-api.git
   cd go-workouts-api
   ```

## 🏃‍♂️ Setup

1. Start the database using Docker Compose:

   ```bash
   docker-compose up -d
   ```

2. Run migrations (handled automatically in `internal/app/app.go`).

3. Build and run the application:
   ```bash
   go build -o main .
   ./main -port 4000
   ```

## 📚 API Endpoints

| Method | Endpoint                 | Description                 |
| ------ | ------------------------ | --------------------------- |
| GET    | `/health`                | Health check                |
| POST   | `/users`                 | Register a new user         |
| POST   | `/tokens/authentication` | Create authentication token |
| GET    | `/workouts/{id}`         | Get workout by ID           |
| POST   | `/workouts`              | Create a new workout        |
| PUT    | `/workouts/{id}`         | Update workout by ID        |
| DELETE | `/workouts/{id}`         | Delete workout by ID        |

All endpoints return JSON responses. Use the token from `/tokens/authentication` for authenticated requests.
