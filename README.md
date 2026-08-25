# Gator

Gator is a command-line RSS feed aggregator written in Go. It stores users, feeds, follows, and posts in PostgreSQL.

## Prerequisites

Install:

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

Make sure PostgreSQL is running and create a database for Gator. The examples below use a local database named `gator_go` with the default `postgres` user:

```sh
createdb gator_go
```

## Install

Install the `gator` CLI with `go install`:

```sh
go install github.com/mumairdotdev/gator-go@latest
```

Go installs the executable in your Go binary directory. Ensure that directory is on your `PATH`, then verify the installation:

```sh
gator
```

You can also run the project directly from a checkout with `go run .`.

## Configure

Create `~/.gatorconfig.json` and set the PostgreSQL connection string. `current_user_name` can be empty until you register a user:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator_go?sslmode=disable",
  "current_user_name": ""
}
```

The application reads this file from your home directory every time it starts. Keep the file private if it contains database credentials.

## Set Up The Database

Run the migrations from the repository root. The project uses [Goose](https://github.com/pressly/goose) migration files:

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator_go?sslmode=disable" up
```

Use the same connection string in `~/.gatorconfig.json` and in the migration command.

## Usage

Register a user. Registration also makes that user the current user in the config file:

```sh
gator register alice
gator users
```

Add and follow an RSS feed:

```sh
gator addfeed "Hacker News" https://news.ycombinator.com/rss
gator feeds
gator following
```

Run the aggregator with the interval between fetches as a Go duration. This command keeps running, so leave it in a separate terminal:

```sh
gator agg 1m
```

After feeds have been collected, browse posts for the current user. The default limit is 2; you can provide a different limit:

```sh
gator browse
gator browse 10
```

Other available commands include:

```text
gator login <username>
gator follow <feed_url>
gator unfollow <feed_url>
gator reset
```

`follow`, `unfollow`, `following`, `browse`, and `addfeed` require a current user. `reset` deletes all users and their associated data, so use it only when you intentionally want to clear the database.
