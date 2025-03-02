package handlers

import (
	"net/http"

	"time"

	"github.com/5h1vanshh/retailer-service/internal/models"
	"github.com/5h1vanshh/retailer-service/internal/repositories"
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate Product ID
	product.ID = "PROD" + time.Now().Format("20060102150405")

	if err := repositories.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product successfully added", "product": product})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateProduct(id, updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Product updated successfully"})
}
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := repositories.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
func GetAllProducts(c *gin.Context) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
