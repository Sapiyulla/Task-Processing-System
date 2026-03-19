# distributed Job Processing Service

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

- API service
- Outbox relay
- Message broker
- Worker service (node)

![](/source/architecture.png)

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

## 7. Дополнительно

> Я решил реализовать проект через `Domain Driven Design` подход. Это значит, что должно быть приложение-документация к каждой доменной модели. 

Вся документация по моделям (_[в папке](/docs/domain/)_):
 - [Task](/docs/domain/task.md)
 - [Execution](/docs/domain/execution.md)