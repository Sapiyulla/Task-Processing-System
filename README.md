# distributed Tasks Processing System

A service for handling various tasks. The peculiarity lies in the distributed nature of the system.

> **THE PROJECT IS EDUCATIONAL**

## Features

- Transactional Outbox pattern
- Idempotent workers
- Message broker integration
- Retry mechanism
- Graceful shutdown

This project demonstrates:

- Transactional Outbox Pattern
- Idempotent Workers
- At-least-once message delivery

## Architecture

The system consists of:

- [API service](/services/api-service/)
- [Outbox relay](/services/relay-outbox/)
- Message broker
- [Worker](/services/worker/)

![architecture](/source/architecture.png)

## Tech Stack

- Go
- PostgreSQL
- Kafka
- Docker

## Потоки данных

1. `User -> API`
2. `API -> DB (task + outbox)`
3. `Relay -> Kafka`
4. `Worker -> Execution`

## 6. Getting Started

### Requirements

- Go 1.22+
- Docker
- PostgreSQL

### Run

```sh
git clone https://github.com/Sapiyulla/Task-Processing-System.git
cd dTPS
docker compose up
```

## Configuration

| Variable |      Description      |
|----------|-----------------------|
|  DB_URL  | PostgreSQL connection |

## Roadmap

- User API
- Task API
- Outbox + Relay
- Worker
