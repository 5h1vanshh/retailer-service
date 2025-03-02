package cmd

import (
	"github.com/5h1vanshh/retailer-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/product", handlers.AddProduct)
	return r
}
