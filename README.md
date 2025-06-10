# Task Queue System

It's a small pet project to learn and master golang and message queues.

## Usage

User sends a request via only API method `api/v1/publish`.
It's a POST method which puts a task in rabbitmq queue.

There are two types of tasks:
    - resize an image
    - send an email

## Installation

To run project you need:

1. Setup `.env` files in `cmd` directories of each services.

```/consumer/cmd/.env
RABBITMQ_URL=amqp://user:pass@rabbitmq:port/

EMAIL_HOST=<provider host like smtp.gmail.com>
EMAIL_PORT=<port of provider, most commonly 587>
EMAIL_USER=<email from message will be sent>
EMAIL_PASS=<pass to email OR app pass like in gmail>
```

```/server/cmd/.env
RABBITMQ_URL=amqp://user:pass@rabbitmq:port/
```

RabbitMQ urls are the same for both services.
`rabbitmq` in host is mandatory because of docker-compose service name

2. Run `docker-compose up --build -d` and enjoy
