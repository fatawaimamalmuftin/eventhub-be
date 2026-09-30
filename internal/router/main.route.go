package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/fatawaimamalmuftin/eventhub-be/docs"
	swagoFiles "github.com/swaggo/files"
	swagoGin "github.com/swaggo/gin-swagger"
)

func MainRouter(r *gin.Engine, db *pgxpool.Pool) {
	r.Use(middleware.Cors)

	r.GET("documentation/*any", swagoGin.WrapHandler(swagoFiles.Handler))

	authRouter(r, db)
	eventRouter(r, db)
}
