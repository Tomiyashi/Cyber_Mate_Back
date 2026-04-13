# Docker Run Guide

## 1) Fill `.env`

Required values:

```env
TELEGRAM_BOT_TOKEN=your_bot_token
POSTGRES_USER=user
POSTGRES_PASSWORD=pass
POSTGRES_DB=cybermate
MINI_APP_URL=https://t.me/CyberMate_bot
```

`DATABASE_URL` is not required for docker compose because it is assembled automatically for the `bot` service.

## 2) Start services

```bash
docker compose up --build -d
```

## 3) Check logs

```bash
docker compose logs -f bot
```

If you see PostgreSQL connected and bot authorized, the app is running correctly.

## 4) Stop services

```bash
docker compose down
```

To remove DB data volume too:

```bash
docker compose down -v
```
