# Blog_aggreator

Gator is a simple CLI that aggregates RSS feeds into a Postgres database and lets you browse posts from the terminal.

## Requirements

- Go (1.20+ recommended)
- PostgreSQL (14+ recommended)

Quick checks:

```bash
go version
psql --version
```

Downloads:

- Go: https://go.dev/dl/
- PostgreSQL: https://www.postgresql.org/download/

Make sure PostgreSQL is running and you have a database created for the app.

## Installation

Install the CLI with `go install`:

```bash
go install github.com/kryptonn36/Blog_aggreator@latest
```

Ensure `$GOPATH/bin` is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Verify:

```bash
gator
```

## Configuration

Gator reads config from a local JSON file that includes the database connection and the current user.

Create a file at:

```text
~/.gatorconfig.json
```

Example:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Do not commit this file; add it to your global gitignore to keep credentials private.

## Database Setup

Gator uses Goose for database migrations.

Install Goose:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Verify:

```bash
goose -version
```

Run migrations:

```bash
make migrate-up
```

Rollback migrations:

```bash
make migrate-down
```

## Usage

Register a user:

```bash
gator register <username>
```

Login:

```bash
gator login <username>
```

Add a feed:

```bash
gator addfeed <feed_url>
```

Follow a feed:

```bash
gator follow <feed_url>
```

Start the aggregator (fetch feeds on an interval):

```bash
gator agg 5s
```

Browse posts:

```bash
gator browse
```

Limit results:

```bash
gator browse 5
```

## Development

Clone the repo and install dependencies (Go modules are handled automatically):

```bash
git clone https://github.com/kryptonn36/Blog_aggreator.git
cd Blog_aggreator
```

Run the CLI directly from source:

```bash
go run . <command>
```

Build a local binary:

```bash
go build -o gator
./gator <command>
```

Run tests:

```bash
go test ./...
```
