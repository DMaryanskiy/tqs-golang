package main

import (
	"github.com/DMaryanskiy/tqs-golang/server/internal/api"
	"github.com/DMaryanskiy/tqs-golang/server/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	config.NewConfig()
	godotenv.Load()

	router := gin.Default()

	router.POST("api/v1/publish", api.PublishTaskHandler)

	router.Run(":8080")
}
