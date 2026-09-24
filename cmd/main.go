package main

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.MainRouter(r)

	r.Run(":1212")
}
