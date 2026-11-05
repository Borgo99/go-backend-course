https://github.com/techschool/simplebank

## Requirements

- [Go](https://go.dev/dl/) 1.27.3 or newer (see `go.mod`)
- [Docker](https://docs.docker.com/get-docker/), to run Postgres locally
- `make`, preinstalled on macOS and Linux
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI, to run the database migrations
- [sqlc](https://sqlc.dev), only needed if you change the SQL queries (the generated code in `db/sqlc` is committed)

Install the CLI tools on macOS:

```bash
brew install golang-migrate sqlc
```

Or with Go on any OS:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Setup

```bash
git clone git@github.com:Borgo99/go-backend-course.git
cd go-backend-course
go mod download

make postgres    # start a Postgres container on port 5432
make createdb    # create the simple_bank database
make migrateup   # create the tables
make test        # run the tests against the database
```

The database runs at `postgresql://root:secret@localhost:5432/simple_bank`. These are local development credentials only.

If the container already exists from an earlier run, start it with `docker start postgres12` instead of `make postgres`.

## Make commands

| Command            | What it does                                     |
| ------------------ | ------------------------------------------------ |
| `make postgres`    | Start the Postgres container                     |
| `make createdb`    | Create the `simple_bank` database                |
| `make dropdb`      | Delete the `simple_bank` database                |
| `make migrateup`   | Apply all migrations                             |
| `make migratedown` | Roll back all migrations                         |
| `make sqlc`        | Regenerate Go code from the queries in `db/query` |
| `make test`        | Run all tests with coverage                      |
