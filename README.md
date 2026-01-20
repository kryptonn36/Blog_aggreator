# Blog_aggreator
📦 Requirements

Before running Gator, make sure you have the following installed:

1. Go (1.20+ recommended)
go version


Install from: https://go.dev/dl/

2. PostgreSQL (14+ recommended)
psql --version


Install from: https://www.postgresql.org/download/

Make sure PostgreSQL is running and you have a database created for the app.

🔧 Installation

Install the gator CLI using go install:

go install github.com/kryptonn36/Blog_aggreator@latest


Ensure $GOPATH/bin is in your PATH:

export PATH=$PATH:$(go env GOPATH)/bin


You should now be able to run:

gator

⚙️ Configuration

Gator uses a config file to store database connection details and the current user.

Create config file

Create a file at:

~/.gatorconfig.json

Example config
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}


<!-- ⚠️ Do not commit this file to GitHub
Add it to .gitignore to keep credentials private. -->

🗄️ Database Setup

Goose (Database Migrations)

Gator uses Goose to manage database migrations.

Install Goose:

go install github.com/pressly/goose/v3/cmd/goose@latest


Verify:

goose -version

Run migrations using Goose:

make migrate-up


To rollback migrations:

make migrate-down

🚀 Usage
Register a user
gator register <username>

Login
gator login <username>

Add a feed
gator addfeed <feed_url>

Follow a feed
gator follow <feed_url>

Start the aggregator (scrape feeds at intervals)
gator agg 5s


(Example: fetch feeds every 5 seconds)

Browse posts
gator browse


Or limit results:

gator browse 5
