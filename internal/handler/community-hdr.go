package handler

import (
	"context"
	"net/http"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	"github.com/gin-gonic/gin"
)

type ICommunitySrv interface {
	GetPopularCommunities(c context.Context) ([]model.PopularCommunity, error)
}

type CommunityHandler struct {
	CS ICommunitySrv
}

func ProviderCommunityHandler(cs ICommunitySrv) *CommunityHandler {
	return &CommunityHandler{
		CS: cs,
	}
}

func (h *CommunityHandler) GetPopularCommunities(c *gin.Context) {
	communities, err := h.CS.GetPopularCommunities(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "popular communities retrieved successfully",
		Data:    communities,
	})
}
