package repositories

import (
	"github.com/5h1vanshh/retailer-service/internal/config"
	"github.com/5h1vanshh/retailer-service/internal/models"
)

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}
func UpdateProduct(id string, updatedFields map[string]interface{}) error {
	return config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updatedFields).Error
}
func GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := config.DB.Where("id = ?", id).First(&product).Error
	return &product, err
}
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}
