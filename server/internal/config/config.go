package config

import "os"

var Cfg *Config

type Config struct {
	RabbitMQUrl string
}

func NewConfig() {
	rabbitMQUrl := os.Getenv("RABBITMQ_URL")
	if rabbitMQUrl == "" {
		rabbitMQUrl = "amqp://guest:guest@localhost:5672/"
	}

	Cfg = &Config{
		RabbitMQUrl: rabbitMQUrl,
	}
}
