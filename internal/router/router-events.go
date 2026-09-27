package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/handler"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func eventRouter(r *gin.Engine, db *pgxpool.Pool) {

	efr := repo.EventsFilterRepo(db)
	efs := service.EventsFilterService(efr)
	efh := handler.EventsFilterHandler(efs)

	r.GET("/events", efh.GetEvents)
}
