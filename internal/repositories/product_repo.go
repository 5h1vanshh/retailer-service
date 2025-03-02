package repositories

import (
	"github.com/5h1vanshh/retailer-service/internal/config"
	"github.com/5h1vanshh/retailer-service/internal/models"
)

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}
