package producer

import (
	"log"

	"github.com/DMaryanskiy/tqs-golang/server/internal/config"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewRabbitMQConnection() *RabbitMQ {
	conf := *config.Cfg
	conn, err := amqp091.Dial(conf.RabbitMQUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel: %s", err)
	}

	_, err = channel.QueueDeclare("tasks", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a RabbitMQ queue: %s", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: channel,
	}
}

func (r *RabbitMQ) PublishTask(queueName string, body []byte) error {
	err := r.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)

	if err != nil {
		log.Fatalf("Failed to publish task: %s", err)
		return err
	}

	return nil
}

func (r *RabbitMQ) CloseConnection() {
	r.Channel.Close()
	r.Conn.Close()
}
