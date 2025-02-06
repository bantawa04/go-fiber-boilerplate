# Go Fiber Boilerplate

A robust Go Fiber boilerplate following the Repository Pattern, designed for building scalable web applications.

## Features

- **Framework**: [Fiber](https://gofiber.io/) - Express-inspired web framework
- **ORM**: [GORM](https://gorm.io/) with PostgreSQL
- **Configuration**: [Viper](https://github.com/spf13/viper) for environment management
- **Validation**: [Go Validator](https://github.com/go-playground/validator)
- **Hot Reload**: [Air](https://github.com/cosmtrek/air) for development
- **Code Generation**: [Gentool](https://gorm.io/gen) for DAO generation
- **Migration tool**: [Migrate](https://github.com/golang-migrate/migrate) for running migration
- **Database**: PostgreSQL
- **Task Runner**: Makefile for common commands
- **Database Transaction**: Middleware-based transaction handling for write operations (Create, Update, Delete) with automatic commit/rollback based on response status
- **Swagger**: API documentation using [Swagger](github.com/gofiber/contrib/swagger)

## Project Structure
| Directory | Description |
|-----------|-------------|
| `app/constants` | Application constants |
| `app/controller` | HTTP request handlers |
| `app/dao` | Data Access Objects |
| `app/dto` | Data Transfer Objects |
| `app/errors` | Custom error definitions |
| `app/middleware` | HTTP middleware |
| `app/model` | Database models |
| `app/repository` | Data access layer |
| `app/request` | Request models |
| `app/response` | Response models |
| `app/service` | Business logic |
| `app/validator` | Request validation |
| `bootstrap` | Application bootstrap |
| `config` | Configuration |
| `database` | Database migrations |
| `docker` | Docker configurations |
| `router` | Route definitions |
| `utils` | Utility functions |


## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/fiber-boilerplate.git
cd fiber-boilerplate
```

2. Set up your environment variables
```bash
cp .env.example .env
```

3. Build the docker image, install dependency and run project
```bash
docker-compose build && docker-compose up -d
```

