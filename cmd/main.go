package main

import (
	"fmt"
	"log"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/config"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("failed to load .env")
		return
	}

	r := gin.Default()

	db, err := config.DBconfig()
	if err != nil {
		fmt.Println(err)
	}

	router.MainRouter(r, db)

	r.Run(":5678")
}
