package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/handler"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func eventRouter(r *gin.Engine, db *pgxpool.Pool) {
	event := r.Group("/event")

	efr := repo.EventsFilterRepo(db)
	efs := service.EventsFilterService(efr)
	efh := handler.EventsFilterHandler(efs)

	event.GET("", efh.GetEvents)

	edr := repo.EventDetailRepo(db)
	eds := service.EventDetailService(edr)
	edh := handler.EventDetailHandler(eds)

	event.GET(":id", edh.GetEventDetail)
}
