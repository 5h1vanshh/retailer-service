package config

import (
	"log"

	"github.com/5h1vanshh/retailer-service/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:RJ04ca5959@5959@tcp(127.0.0.1:3306)/retailer_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db

	// Auto-migrate tables
	err = DB.AutoMigrate(&models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully.")
}
