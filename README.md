# go-crm-fiber

A small CRM backend written in Go using the Fiber web framework. This project implements basic CRUD operations for Leads and provides a lightweight starting point for building a contacts/leads management API.

Note: This README was written from the repository structure. It assumes the project exposes a basic Leads CRUD API via Fiber. If your implementation differs (routes or DB), adjust the examples below.

## Table of contents

- [Features](#features)
- [Requirements](#requirements)
- [Quick start (Windows - PowerShell)](#quick-start-windows---powershell)
- [Project structure](#project-structure)
- [Environment variables](#environment-variables)
- [API (assumed)](#api-assumed)
- [Development notes](#development-notes)
- [License](#license)

## Features

- Simple REST API for managing Leads (create, read, update, delete)
- Uses the Fiber web framework
- Simple database initialization in `database/database.go`

## Requirements

- Go 1.18+ installed and available in your PATH
- Git (optional, for cloning)
- A supported database if the project expects one (see `database/database.go`). The project may default to an embedded or local DB depending on its implementation.

## Quick start (Windows - PowerShell)

Open PowerShell in the `go-crm-fiber` directory and run:

```powershell
# install dependencies and tidy modules
go mod tidy

# run the app
go run main.go

# or build and run
go build -o crm.exe
./crm.exe
```

The server will typically listen on a port defined in the `.env` file or `main.go` (commonly `:3000` or `:8080`). Watch the console output for the actual listen address.

## Project structure

- `main.go` - application entrypoint and route wiring
- `.env` - environment variables used by the app (not committed)
- `go.mod` - Go module file
- `database/database.go` - database setup and connection
- `lead/lead.go` - lead model and handlers (CRUD)

## Environment variables

Place runtime configuration in the `.env` file at the project root. Common variables the app may expect:

- `PORT` - port to run the server on (e.g. `3000`)
- `DATABASE_URL` or other DB-specific variables - connection string for your database

Example `.env` (adjust to your DB and needs):

```
PORT=3000
DATABASE_URL=sqlite3://db.sqlite
```

If the project uses a different DB (Postgres, MySQL), set the correct connection string and ensure the DB is running.

## API (assumed endpoints)

The repository likely implements a Leads resource. Example endpoints (adjust to your actual routes):

- GET /leads - list all leads
- GET /leads/:id - get a single lead by id
- POST /leads - create a new lead (JSON body)
- PUT /leads/:id - update an existing lead (JSON body)
- DELETE /leads/:id - delete a lead

Example curl (PowerShell):

```powershell
# Create a lead
curl -X POST http://localhost:3000/leads -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@example.com"}'

# List leads
curl http://localhost:3000/leads

# Get lead with id 1
curl http://localhost:3000/leads/1

# Update lead
curl -X PUT http://localhost:3000/leads/1 -H "Content-Type: application/json" -d '{"name":"Alice A.","email":"alice@newdomain.com"}'

# Delete lead
curl -X DELETE http://localhost:3000/leads/1
```

If your project registers different route prefixes (for example `/api/leads`) or different field names, update the examples accordingly.

## Development notes

- To add new routes, update `main.go` and the handler functions under `lead/lead.go`.
- To change the DB, edit `database/database.go` and update connection strings in `.env`.
- Run `go vet` and `golangci-lint` (if installed) to check for issues.

## Contributing

Feel free to open an issue or a pull request. For small changes, please follow the existing code style and test locally before submitting.

## License

This repository does not include a license file. Add a `LICENSE` file if you want to make the project open-source with a specific license.

---

If you want, I can:

- inspect the source files and update the README with exact endpoints and DB instructions
- add sample Postman collection or simple tests
- scaffold a Dockerfile / docker-compose for easy local DB + server startup

Tell me which of those you'd like next.
