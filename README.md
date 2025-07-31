# telegram-bot-go-starter
[อ่านคู่มือภาษาไทย](https://medium.com/@nipatchapakdee/สร้าง-telegram-bot-ด้วย-go-สำหรับมือใหม่-cbdf4625c780)
Starter project for building a Telegram Bot using Go, PostgreSQL, and Docker Compose.

> Designed for beginners who want to quickly start building and deploying a Telegram bot connected to a PostgreSQL database.

---

## Tech Stack

- **Go** – for building the bot and backend services
- **PostgreSQL** – for data storage
- **Docker Compose** – for easy development setup
- **Swagger** – for documenting and testing APIs

---

## Features

- Connects with Telegram Bot via `TELEGRAM_TOKEN`
- Uses PostgreSQL for persistent data storage
- Runs fully in Docker (Postgres + pgAdmin)
- Swagger docs for API testing and development
- Ready-to-extend structure for services, controllers, and database logic

---

## Getting Started

### 1. Clone this repository

```bash
git clone https://github.com/Peenipat/telegram-bot-go-starter.git

cd telegram-bot-go-starter
```


### 2. Create .env file
Create a .env file in the root directory with the following variables

```bash
TELEGRAM_TOKEN=your-telegram-bot-token
DB_URL=postgres://<user>:<password>@<host>:<port>/<database>?sslmode=disable
```

### 3. Start the database using Docker Compose
```bash
docker compose up -d
```

This will start
PostgreSQL on port 5433

pgAdmin on port 5050

You can access pgAdmin at http://localhost:5050
Login credentials

```bash
Email: admin@admin.com
Password: admin
```

### 4. Run the Go Application

```bash
go run main.go
```

### 5. Access Swagger API Documentation

```bash
http://localhost:8080/swagger/index.html
```
