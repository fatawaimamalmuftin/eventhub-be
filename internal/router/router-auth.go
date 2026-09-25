package router

import (
	"github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRouter(r *gin.Engine, db *pgxpool.Pool) {
	auth := r.Group("/auth")

	auth.GET("regis", func(c *gin.Context) {
		if e := c.ShouldBindBodyWithJSON(&dto.Regis{}); e != nil {
			c.Error(cuserror.ErrBinding)
			return
		}
	})
}
