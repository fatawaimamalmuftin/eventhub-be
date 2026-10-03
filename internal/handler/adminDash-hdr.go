package handler

import (
	"context"
	"net/http"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/gin-gonic/gin"
)

type IAdminDashHdr interface {
	GetAdminDashSrv(c context.Context) (*dto.AdminDashboardResponse, error)
}

type AdminDashHdr struct {
	ADh IAdminDashHdr
}

func ProviderAdminDashHandler(adh IAdminDashHdr) *AdminDashHdr {
	return &AdminDashHdr{
		ADh: adh,
	}
}

func (pr *AdminDashHdr) GetAdminDashHdr(c *gin.Context) {
	data, err := pr.ADh.GetAdminDashSrv(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: cuserror.InternalError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}
