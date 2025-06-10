package producer_test

import (
	"os"
	"testing"

	"github.com/DMaryanskiy/tqs-golang/server/internal/config"
	"github.com/DMaryanskiy/tqs-golang/server/internal/producer"
	"github.com/stretchr/testify/assert"
)

func init() {
	_ = os.Setenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	config.NewConfig()
}

func TestNewRabbitMQConnection(t *testing.T) {
	rmq := producer.NewRabbitMQConnection()
	assert.NotNil(t, rmq.Conn, "Expected connection to be initialized")
	assert.NotNil(t, rmq.Channel, "Expected channel to be initialized")
	rmq.CloseConnection()
}

func TestPublishTask(t *testing.T) {
	rmq := producer.NewRabbitMQConnection()
	defer rmq.CloseConnection()

	body := []byte(`{"type":"dummy_task","payload":{"test":"value"}}`)
	err := rmq.PublishTask("tasks", body)

	assert.NoError(t, err, "Expected task to be published successfully")
}
