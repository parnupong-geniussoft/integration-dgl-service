# integration-dgl-service

Dgl service written in Go (Golang) following the Clean Architecture approach.

## 🔧 Project Structure

```
.
│   .env
│   .env.debug.example
│   .env.localhost.example
│   docker-compose.yml
│   Dockerfile
│   go.mod
│   go.sum
│   README.md
│
├───.github
│   └───workflows
│           build.yml
│
├───.vscode
│       launch.json
│
├───app
│       main.go
│
├───configs
│       config.go
│
├───docs
│       docs.go
│       swagger.json
│       swagger.yaml
│
├───modules
│   ├───dgl
│   │   ├───controllers
│   │   │       dgl_controllers.go
│   │   │
│   │   ├───entities
│   │   │       dgl_entities.go
│   │   │       master_entities.go
│   │   │
│   │   ├───repositories
│   │   │       dql_repositories.go
│   │   │
│   │   └───usecases
│   │           dql_usecases.go
│   │
│   ├───middlewares
│   │       basic_auth.go
│   │       jwt.go
│   │       loggers.go
│   │       recover.go
│   │
│   └───servers
│           handler.go
│           server.go
│
└───pkg
    ├───databases
    │       postgresql.go
    │
    ├───loggers
    │       loggers.go
    │       loggers_entities.go
    │       loggers_resty.go
    │
    ├───resty_factory
    │       resty_factory.go
    │
    └───utils
            array.go
            connection_url_builder.go
            hash.go
            time.go
```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/parnupong-geniussoft/integration-dgl-service.git
cd integration-dgl-service
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Generate API documentation (Swagger)

```bash
swag init --generalInfo app/main.go 
```

### 4. Run the Project

```bash
go run ./app/main.go
```

## 📌 Features

- ✅ Submit Dgl list to Ktb

## 🛠 Technologies Used

- [Go Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Swagger (Swaggo)](https://github.com/swaggo/swag)

## 📄 API Usage

### POST /dql

## 📚 API Documentation

Swagger UI available at:  
[http://localhost:8082/swagger/index.html](http://localhost:8082/swagger/index.html)

## 🧹 Clean Architecture Overview

- `controllers` → Handles requests and invokes usecases
- `usecases` → Business logic implementation
- `repositories` → Handles data persistence
- `entities` → Data structures

## 📁 Middleware

- `middlewares/basic_auth.go` → Basic authentication verification
- `middlewares/jwt.go` → JWT Token verification
- `middlewares/loggers.go` → Logs request data into database
- `middlewares/recover.go` → Handles panic and returns safe response
