package handlers

import (
	"net/http"
	"time"

	"github.com/5h1vanshh/retailer-service/internal/models"
	"github.com/5h1vanshh/retailer-service/internal/repositories"
	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowed, err := repositories.CheckLastOrderTime(order.CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking cooldown"})
		return
	}

	if !allowed {
		c.JSON(http.StatusForbidden, gin.H{"error": "You must wait 5 minutes before placing another order"})
		return
	}

	order.ID = "ORD" + time.Now().Format("20060102150405") // Generate a simple unique ID
	order.Status = "order placed"
	order.CreatedAt = time.Now()

	if err := repositories.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not place order"})
		return
	}

	c.JSON(http.StatusOK, order)
}
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := repositories.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
