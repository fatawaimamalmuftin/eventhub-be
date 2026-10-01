package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	_ "github.com/fatawaimamalmuftin/eventhub-be/docs"
	swagoFiles "github.com/swaggo/files"
	swagoGin "github.com/swaggo/gin-swagger"
)

func MainRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	r.Use(middleware.Cors)

	r.GET("documentation/*any", swagoGin.WrapHandler(swagoFiles.Handler))

	authRouter(r, db)
	eventRouter(r, db, rdb)
}
