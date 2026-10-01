package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/handler"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/middleware"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func userRouter(r *gin.Engine, db *pgxpool.Pool) {
	user := r.Group("/user")

	jer := repo.JoinEventRepo(db)
	jes := service.JoinEventService(jer)
	jeh := handler.JoinEventHandler(jes)

	user.GET(":eventId/join", middleware.CheckJWTtoken, jeh.CreateJoinEventHandler)

	ler := repo.LeaveEventRepo(db)
	les := service.LeaveEventService(ler)
	leh := handler.LeaveEventHandler(les)

	user.GET(":eventId/leave", middleware.CheckJWTtoken, leh.CreateLeaveEventHandler)
}
