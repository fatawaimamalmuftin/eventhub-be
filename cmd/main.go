package main

import (
	"fmt"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/config"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := config.DBconfig()
	if err != nil {
		fmt.Println(err)
	}

	r.GET("/ping")

	router.MainRouter(r, db)

	r.Run(":5678")
}
