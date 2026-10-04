package handler

import (
	"context"
	"net/http"
	"strconv"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	"github.com/gin-gonic/gin"
)

type ICommunityDetailSrv interface {
	GetCommunityDetail(id int, c context.Context) (model.CommunityDetail, error)
}

type CommunityDetailHandler struct {
	CDh ICommunityDetailSrv
}

func ProviderCommunityDetailHandler(cdh ICommunityDetailSrv) *CommunityDetailHandler {
	return &CommunityDetailHandler{
		CDh: cdh,
	}
}

func (h *CommunityDetailHandler) GetCommunityDetail(c *gin.Context) {
	communityID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "invalid community id",
		})
		return
	}

	community, err := h.CDh.GetCommunityDetail(
		communityID,
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: cuserror.InternalError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "community detail retrieved successfully",
		Data:    community,
	})
}
