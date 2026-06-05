# Concave Tech

A layered Go REST API that serves user data from PostgreSQL using the Gin framework.

## Overview

This repository contains a small service-oriented backend organized into clear layers:

- **Controller layer** (`controllers/`): HTTP route registration and request/response handling.
- **Service layer** (`services/`): Business logic and DTO mapping.
- **Repository layer** (`repositories/`): SQL queries and persistence access.
- **Entity/DTO models** (`entities/`, `dtos/`): Internal and API-facing data contracts.
- **Infrastructure utilities** (`utilities/`): Database connection bootstrap.

Main entrypoint: `main.go`

## Tech Stack

- Go 1.24+
- Gin (`github.com/gin-gonic/gin`)
- PostgreSQL (`github.com/lib/pq`)
- Docker Compose (local PostgreSQL + pgAdmin)

## API

### `GET /users`

Returns all users from PostgreSQL.

Example response:

```json
[
  {
    "firstname": "John",
    "lastname": "Citizen",
    "date_of_birth": "03-Mar-2003",
    "income": 65000
  }
]
```

## Prerequisites

- Go installed
- PostgreSQL running (locally or via Docker)

## Database Setup

The API queries the table `"user"` with columns:

- `id`
- `firstname`
- `lastname`
- `date_of_birth`
- `income`

Use the schema below if you are initializing the DB manually:

```sql
CREATE TABLE IF NOT EXISTS "user" (
  id SERIAL PRIMARY KEY,
  firstname VARCHAR(50) NOT NULL,
  lastname VARCHAR(50) NOT NULL,
  date_of_birth DATE NOT NULL,
  income NUMERIC(12,2) NOT NULL
);
```

Optional seed data:

```sql
INSERT INTO "user" (firstname, lastname, date_of_birth, income)
VALUES ('John', 'Citizen', '2003-03-03', 65000);
```

## Running with Docker Compose (Database Only)

From the repository root:

```bash
docker compose up -d
```

This starts:

- PostgreSQL on `localhost:5432`
- pgAdmin on `http://localhost:8858`

## Running the API Locally

```bash
go mod download
go run main.go
```

The API starts on `http://localhost:8080`.

## Configuration

The DB host is controlled by `PGHOST`:

- Default: `localhost`
- Example for container networking: `PGHOST=postgres`

Connection string credentials are currently defined in `utilities/db.go`.
Update them to match your PostgreSQL credentials before running in your environment.

## Validation

```bash
go test ./...
```

> Note: there is an additional sample Go program (`person.go`) in the repository root that also declares a `main` function. Running tests/builds across all root package files may fail unless that sample is excluded or moved.

## Project Structure

```text
.
├── controllers/
├── dtos/
├── entities/
├── repositories/
├── services/
├── utilities/
├── main.go
├── docker-compose.yaml
├── create.sql
└── data.sql
```
