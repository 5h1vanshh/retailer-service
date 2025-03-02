package repositories

import (
	"time"

	"github.com/5h1vanshh/retailer-service/internal/config"
	"github.com/5h1vanshh/retailer-service/internal/models"
)

// CheckLastOrderTime checks if the customer is allowed to place an order (5 min cooldown)
func CheckLastOrderTime(customerID string) (bool, error) {
	var lastOrder models.Order
	err := config.DB.Where("customer_id = ?").Order("created_at desc").First(&lastOrder).Error
	if err != nil {
		return true, nil // No previous orders
	}

	timeSinceLastOrder := time.Since(lastOrder.CreatedAt)
	return timeSinceLastOrder >= 5*time.Minute, nil
}

// CreateOrder inserts a new order into the database
func CreateOrder(order *models.Order) error {
	return config.DB.Create(order).Error
}
func GetOrderByID(id string) (*models.Order, error) {
	var order models.Order
	err := config.DB.Where("id = ?", id).First(&order).Error
	return &order, err
}
