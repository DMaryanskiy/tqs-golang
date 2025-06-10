package consumer_test

import (
	"os"
	"testing"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/config"
	"github.com/DMaryanskiy/tqs-golang/consumer/internal/consumer"
)

func init() {
	_ = os.Setenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	config.NewConfig()
}


// For this test you need to run RabbitMQ via docker manually
func TestNewRabbitMQConnection(t *testing.T) {
	rmq := consumer.NewRabbitMQConnection()
	if rmq.Conn == nil {
		t.Fatal("Expected RabbitMQ connection, got nil")
	}
	if rmq.Channel == nil {
		t.Fatal("Expected RabbitMQ channel, got nil")
	}
	rmq.CloseConnection()
}

func TestQueueDeclared(t *testing.T) {
	rmq := consumer.NewRabbitMQConnection()
	defer rmq.CloseConnection()

	_, err := rmq.Channel.QueueDeclarePassive(
		"tasks", true, false, false, false, nil,
	)
	if err != nil {
		t.Fatalf("Expected 'tasks' queue to be declared, but got error: %s", err)
	}
}
