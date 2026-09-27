package handler

import (
	"net/http"
	"strings"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/gin-gonic/gin"
)

type ILogoutSrv interface {
	Logout(token string) error
}

type LogoutHdr struct {
	Ls ILogoutSrv
}

func LogoutHandler(logoutService ILogoutSrv) *LogoutHdr {
	return &LogoutHdr{
		Ls: logoutService,
	}
}

func (l *LogoutHdr) Logout(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	bearerPath := strings.Split(authorization, " ")

	if len(bearerPath) != 2 {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid bearer token",
		})
		return
	}

	if err := l.Ls.Logout(bearerPath[1]); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "logout success",
	})
}
