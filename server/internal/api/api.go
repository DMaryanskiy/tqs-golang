package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DMaryanskiy/tqs-golang/server/internal/producer"
	"github.com/DMaryanskiy/tqs-golang/server/internal/tasks"
	"github.com/gin-gonic/gin"
)

func PublishTaskHandler(c *gin.Context) {
	var newTask tasks.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	rabbitmqQueue := "tasks"
	body, err := json.Marshal(newTask)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}

	r := producer.NewRabbitMQConnection()
	defer r.CloseConnection()

	r.PublishTask(rabbitmqQueue, body)

	c.JSON(http.StatusOK, gin.H{"message": "Task was sent to RabbitMQ queue"})
}
