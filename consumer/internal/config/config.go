package config

import "os"

var Cfg *Config

type Config struct {
	RabbitMQUrl string

	EmailHost string
	EmailPort string
	EmailUser string
	EmailPass string
}

func NewConfig() {
	rabbitMQUrl := os.Getenv("RABBITMQ_URL")
	if rabbitMQUrl == "" {
		rabbitMQUrl = "amqp://guest:guest@localhost:5672/"
	}
	
	emailHost := os.Getenv("EMAIL_HOST")
	if emailHost == "" {
		emailHost = "smtp://localhost"
	}

	emailPort := os.Getenv("EMAIL_PORT")
	if emailPort == "" {
		emailPort = "587"
	}

	emailUser := os.Getenv("EMAIL_USER")
	if emailUser == "" {
		emailUser = "guest"
	}

	emailPass := os.Getenv("EMAIL_PASS")
	if emailPass == "" {
		emailPass = "guest"
	}

	Cfg = &Config{
		RabbitMQUrl: rabbitMQUrl,

		EmailHost: emailHost,
		EmailPort: emailPort,
		EmailUser: emailUser,
		EmailPass: emailPass,
	}
}
