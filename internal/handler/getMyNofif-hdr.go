package handler

import (
	"context"
	"net/http"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	jwtpkg "github.com/fatawaimamalmuftin/eventhub-be/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type IGetMyNofif interface {
	GetMyNotifSrv(userId int, c context.Context) ([]model.MyNotif, error)
}

type GetMyNotifHandler struct {
	GMs IGetMyNofif
}

func ProviderGetMyNotifHandler(gms IGetMyNofif) *GetMyNotifHandler {
	return &GetMyNotifHandler{
		GMs: gms,
	}
}

func (pr *GetMyNotifHandler) GetMyNotifHdr(c *gin.Context) {
	tokenClaims, ok := c.Get("tokenCleims")
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "please login first",
		})
		return
	}

	token, ok := tokenClaims.(jwtpkg.JWTclem)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid token",
		})
	}

	data, err := pr.GMs.GetMyNotifSrv(token.Id, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: cuserror.InternalError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "my notif retrieved successfully",
		Data:    data,
	})
}
