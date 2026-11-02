# Book Rest Backend

A REST API for managing books with JWT-based user authentication, built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io) following a clean, layered architecture (handler → usecase → repository).

## Features

- **User Authentication**: User registration and login with password hashing and JWT-based authentication.
- **Book Management (CRUD)**: Endpoints to create, view, update, and delete books.
- **Pagination**: Paginated book listings with configurable `page` and `limit` query parameters.
- **Request Validation**: Automatic request payload validation with structured error responses.
- **Interactive API Documentation**: Swagger/OpenAPI documentation served through Swagger UI.
- **Database Migrations**: Version-controlled database schema management using `golang-migrate`.
- **Clean Architecture**: Decoupled layered design (Handler → Usecase → Repository).
- **Docker Support**: Docker and Docker Compose configuration for quick setup and deployment.

## Tech Stack

- **Language**: [Go](https://go.dev/) (1.26+)
- **HTTP Web Framework**: [Gin](https://github.com/gin-gonic/gin) (v1.12.0)
- **ORM**: [GORM](https://gorm.io)
- **JWT Authentication**: [golang-jwt](https://github.com/golang-jwt/jwt)
- **Request Validation**: [validator](https://github.com/go-playground/validator)
- **Swagger/OpenAPI Documentation**: [swaggo](https://github.com/swaggo/swag)
- **Testing**: [testify](https://github.com/stretchr/testify)
- **Database Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Environment Variable**: [godotenv](https://github.com/joho/godotenv)

## Project Structure

```text
├── assets/                 # Static assets
├── cmd/                    # Application entry point
├── database/
│   └── migrations/         # Database schema migrations
├── docs/                   # Generated Swagger/OpenAPI documentation
├── internal/
│   ├── config/             # Application and service configuration
│   ├── delivery/
│   │   ├── handler/        # HTTP request handlers
│   │   ├── middleware/     # HTTP middleware
│   │   ├── request/        # Request payload definitions
│   │   ├── response/       # API response definitions
│   │   └── routes/         # API route registration
│   ├── entity/             # Domain entities
│   ├── helper/             # Shared helpers
│   ├── model/              # Database models
│   ├── repository/         # Data access layer
│   └── usecase/
│       ├── book/
│       │   └── mock/       # Book use-case mocks
│       └── user/
│           └── mock/       # User use-case mocks
└── test/                   # API tests
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.26+
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/) (for running MySQL)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate)
- [swag CLI](https://github.com/swaggo/swag) (for generating API docs)
- [godotenv CLI](https://github.com/joho/godotenv)

## Environment Variables

Create a `.env` file in the project root with the following variables:

```env
APP_NAME=
APP_HOST=
APP_PORT=

SWAGGO_PORT=

AUTH_TOKEN_DURATION=
AUTH_SIGNATURE_KEY=
AUTH_CONTEXT_KEY=

DATABASE_HOST=
DATABASE_PORT=
DATABASE_USER=
DATABASE_PASSWORD=
DATABASE_ROOT_PASSWORD=
DATABASE_NAME=
```

## Getting Started

1. Clone the repository and install dependencies:

   ```bash
   go mod download
   ```

2. Start the MySQL database:

   ```bash
   docker-compose up -d
   ```

3. Run database migrations:

   ```bash
   make migration-up
   ```

4. Run the API server:

   ```bash
   make run-api
   ```

The server will start on the host/port configured in `.env` (`APP_HOST`/`APP_PORT`).

## API Documentation

Swagger docs are generated with `swaggo` and served via `gin-swagger`. After starting the server, visit:

```
http://<APP_HOST>:<APP_PORT>/api/v1/docs/index.html
```

To regenerate the docs after changing handlers/annotations:

```bash
make api-docs
```

## API Endpoints

| Method | Endpoint          | Description                  | Auth Required |
|--------|-------------------|-------------------------------|----------------|
| GET    | `/`               | Health check                 | No             |
| GET    | `/robots.txt`     | Crawler policy               | No             |
| POST   | `/users/register` | Register a new user          | No             |
| POST   | `/users/login`    | Login and receive a JWT      | No             |
| GET    | `/users/current`  | Get the current authenticated user | Yes      |
| POST   | `/books/`         | Create a book                 | Yes            |
| GET    | `/books/`         | List all books                | Yes            |
| GET    | `/books/:id`      | Get a book by ID              | Yes            |
| PUT    | `/books/:id`      | Update a book by ID           | Yes            |
| DELETE | `/books/:id`      | Delete a book by ID           | Yes            |

## Database Migrations

Create a new migration:

```bash
make migration
```

Apply migrations:

```bash
make migration-up
```

Rollback migrations:

```bash
make migration-down
```

## Testing

Run the test suite:

```bash
make run-test
```

## License

This project is licensed under the [MIT License](LICENSE).