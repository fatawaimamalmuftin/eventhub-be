package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/handler"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func authRouter(r *gin.Engine, db *pgxpool.Pool) {

	Auth := r.Group("/auth")

	rr := repo.RegisRepo(db)
	rs := service.RegisService(rr)
	rh := handler.RegisHandler(rs)
	Auth.POST("regis", rh.Regis)

	lr := repo.LoginRepo(db)
	ls := service.LoginService(lr)
	lh := handler.LoginHandler(ls)
	Auth.POST("login", lh.Login)
}
