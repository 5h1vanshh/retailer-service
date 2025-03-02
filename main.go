package main

import (
	"github.com/5h1vanshh/retailer-service/cmd"
	"github.com/5h1vanshh/retailer-service/internal/config"
)

func main() {
	config.ConnectDB()
	r := cmd.SetupRouter()
	r.Run(":8080")
}
