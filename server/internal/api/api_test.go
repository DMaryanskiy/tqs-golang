package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DMaryanskiy/tqs-golang/server/internal/api"
	"github.com/DMaryanskiy/tqs-golang/server/internal/config"
	"github.com/DMaryanskiy/tqs-golang/server/internal/producer"
	"github.com/DMaryanskiy/tqs-golang/server/internal/tasks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	_ = os.Setenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	config.NewConfig()
	producer.NewRabbitMQConnection()
}

func TestPublishTaskHandler_ValidTask(t *testing.T) {
	router := gin.Default()
	router.POST("/publish", api.PublishTaskHandler)

	task := tasks.Task{
		Type: "dummy_task",
		Payload: json.RawMessage(`{
			"some_data": "test"
		}`),
	}
	body, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/publish", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task was sent")
}

func TestPublishTaskHandler_InvalidJSON(t *testing.T) {
	router := gin.Default()
	router.POST("/publish", api.PublishTaskHandler)

	invalidJSON := []byte(`{ invalid json }`)
	req, _ := http.NewRequest("POST", "/publish", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid request body")
}
