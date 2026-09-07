# Book Rest Backend

A REST API for managing books with JWT-based user authentication, built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io) following a clean, layered architecture (handler → usecase → repository).

## Tech Stack

- [Go](https://go.dev/) 1.26
- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [GORM](https://gorm.io) + MySQL driver — ORM
- [golang-jwt](https://github.com/golang-jwt/jwt) — JWT authentication
- [validator](https://github.com/go-playground/validator) — request validation
- [swaggo](https://github.com/swaggo/swag) — Swagger/OpenAPI documentation
- [testify](https://github.com/stretchr/testify) — testing
- [golang-migrate](https://github.com/golang-migrate/migrate) — database migrations
- [godotenv](https://github.com/joho/godotenv) — environment variable loading

## Project Structure

```
cmd/                  Application entrypoint
database/migrations/  SQL migration files
docs/                 Generated Swagger docs
internal/
  config/             Env, database, Gin, and Swagger bootstrap
  delivery/
    handler/           HTTP request handlers
    middleware/         Gin middlewares (e.g. auth)
    request/            Request payload structs
    response/           Response payload structs
  entity/               Domain entities
  helper/                Shared helpers (e.g. auth/JWT)
  model/                 Database models
  repository/            Data access layer
  usecase/               Business logic layer (+ mocks for testing)
test/                  Integration/unit tests
```

## Prerequisites

- Go 1.26+
- Docker & Docker Compose (for MySQL)
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
http://localhost:<APP_PORT>/swagger/index.html
```

To regenerate the docs after changing handlers/annotations:

```bash
make api-docs
```

## API Endpoints

| Method | Endpoint          | Description                  | Auth Required |
|--------|-------------------|-------------------------------|----------------|
| POST   | `/users/register` | Register a new user          | No             |
| POST   | `/users/login`    | Login and receive a JWT      | No             |
| GET    | `/users/current`  | Get the current authenticated user | Yes      |
| POST   | `/books/`         | Create a book                 | -              |
| GET    | `/books/`         | List all books                | -              |
| GET    | `/books/:id`      | Get a book by ID              | -              |
| PUT    | `/books/:id`      | Update a book by ID           | -              |
| DELETE | `/books/:id`      | Delete a book by ID           | -              |

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

This project is licensed under the terms of the [LICENSE](LICENSE) file.