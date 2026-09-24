package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MainRouter(r *gin.Engine){
	r.Use(middleware.Cors)
}