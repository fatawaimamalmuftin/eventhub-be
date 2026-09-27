package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MainRouter(r *gin.Engine, db *pgxpool.Pool) {
	r.Use(middleware.Cors)

	authRouter(r, db)
	eventRouter(r, db)
}
