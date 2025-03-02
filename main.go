package main

import (
	"github.com/5h1vanshh/retailer-service/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.Run(":8080")
}
