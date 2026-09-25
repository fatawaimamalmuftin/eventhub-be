package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MainRouter(r *gin.Engine, pool *pgxpool.Pool) {
	r.Use(middleware.Cors)

	authRouter(r, pool)
}
