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

## Project Structure
.
├── app/
│   ├── constants/     # Application constants
│   ├── controller/    # HTTP request handlers
│   ├── dao/          # Data Access Objects
│   ├── dto/          # Data Transfer Objects
│   ├── errors/       # Custom error definitions
│   ├── middleware/   # HTTP middleware
│   ├── model/        # Database models
│   ├── repository/   # Data access layer
│   ├── request/      # Request models
│   ├── response/     # Response models
│   ├── service/      # Business logic
│   └── validator/    # Request validation
├── bootstrap/        # Application bootstrap
├── config/          # Configuration
├── database/        # Database migrations
├── docker/         # Docker configurations
├── router/         # Route definitions
└── utils/          # Utility functions


## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/fiber-boilerplate.git
cd fiber-boilerplate