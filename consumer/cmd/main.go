package main

import (
	"encoding/json"
	"log"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/config"
	"github.com/DMaryanskiy/tqs-golang/consumer/internal/consumer"
	"github.com/DMaryanskiy/tqs-golang/consumer/internal/tasks"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.NewConfig()

	r := consumer.NewRabbitMQConnection()
	defer r.CloseConnection()

	msgs, err := r.Channel.Consume("tasks", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a RabbitMQ consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var t tasks.Task
			err := json.Unmarshal(msg.Body, &t)
			if err != nil {
				log.Printf("Invalid task format: %s", err)
				continue
			}
			tasks.HandleTask(t)
		}
	} ()

	log.Println("Waiting for tasks. To exit precc Ctrl+C")
	<-forever
}
