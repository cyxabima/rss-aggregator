# GopherFeed: RSS Aggregator & Scraper

A high-performance RSS feed aggregator and microservice built in **Go**. This project allows users to register, add their favorite RSS feeds, follow/unfollow them, and automatically aggregate posts into a centralized database via a concurrent background worker.

## Features

* **User Management**: API key-based authentication for secure access.
* **Feed Management**: Add and manage multiple RSS feeds (XML) for tracking.
* **Following System**: "Multiplayer" functionality to follow/unfollow feeds added by other users.
* **Concurrent Scraper**: A robust background worker that uses Go routines to fetch and parse feeds at a configurable interval.
* **Database Persistence**: Fully typed database interactions using PostgreSQL.

## Tech Stack

* **Language**: [Go (Golang)](https://golang.org/)
* **Database**: [PostgreSQL](https://www.postgresql.org/)
* **Database Tooling**:
* [SQLC](https://sqlc.dev/) (Generate type-safe Go code from SQL)
* [Goose](https://github.com/pressly/goose) (Database migrations)


* **Frameworks/Libs**:
* `chi` (Router)
* `godotenv` (Configuration)
* `google/uuid` (ID generation)



## Installation & Setup

### Prerequisites

* Go 1.20+ installed
* PostgreSQL installed and running
* `sqlc` and `goose` binaries installed:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest

```



### Steps

1. **Clone the repo**:
```bash
git clone https://github.com/your-username/rss-aggregator.git
cd rss-aggregator

```


2. **Environment Configuration**:
Create a `.env` file in the root directory:
```env
PORT=8080
DB_URL=postgres://user:password@localhost:5432/rssagg?sslmode=disable

```


3. **Database Migrations**:
Navigate to your schema directory and run migrations:
```bash
cd sql/schema
goose postgres <YOUR_DB_URL> up

```


4. **Build and Run**:
```bash
go build -o rssagg
./rssagg

```



## API Endpoints

| Method | Endpoint           | Description                          | Auth Required |
| ------ | ------------------ | ------------------------------------ | ------------- |
| `GET`  | `/v1/healthz`      | Check API health                     | No            |
| `POST` | `/v1/users`        | Register a new user                  | No            |
| `GET`  | `/v1/users`        | Get current user info                | Yes           |
| `POST` | `/v1/feeds`        | Add a new RSS feed                   | Yes           |
| `GET`  | `/v1/feeds`        | Get all registered feeds             | No            |
| `POST` | `/v1/feed_follows` | Follow a feed                        | Yes           |
| `GET`  | `/v1/posts`        | Get latest posts from followed feeds | Yes           |

## What I Learned

* **SQLC vs ORMs**: Writing raw SQL for performance while maintaining type safety in Go.
* **Concurrency**: Using `time.Ticker` and Go routines to build a background worker that scrapes data without blocking the main API.
* **Middleware**: Implementing custom authentication middleware to protect sensitive routes.
* **Environment Management**: Handling production-grade configurations using `.env` files.

---

*Developed with the help of the [Boot.dev](https://www.boot.dev) Backend Development guided project*

