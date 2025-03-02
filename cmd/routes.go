package cmd

import (
	"github.com/5h1vanshh/retailer-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/product", handlers.AddProduct)
	r.PATCH("/product/:id", handlers.UpdateProduct)
	r.GET("/product/:id", handlers.GetProductByID)
	r.GET("/products", handlers.GetAllProducts)
	r.POST("/order", handlers.PlaceOrder)
	r.GET("/order/:id", handlers.GetOrderByID)

	return r
}
