package handlers

import (
	"net/http"

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

	if err := repositories.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product successfully added", "product": product})
}
