package main

import (
	"log"

	"github.com/5h1vanshh/retailer-service/cmd"
	"github.com/5h1vanshh/retailer-service/internal/config"
)

func main() {
	config.ConnectDB()

	r := cmd.SetupRouter()
	log.Println("Server is running on http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
