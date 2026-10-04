package handler

import (
	"context"
	"net/http"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/gin-gonic/gin"
)

type ICreateEvent interface {
	CreateEventService(c context.Context, data dto.CreateEvent) (string, error)
}

type CreateEventHandler struct {
	CEs ICreateEvent
}

func ProviderCreateEventHandler(ces ICreateEvent) *CreateEventHandler {
	return &CreateEventHandler{
		CEs: ces,
	}
}

func (pr *CreateEventHandler) CreateEventHdr(c *gin.Context) {
	var event dto.CreateEvent

	if e := c.ShouldBindBodyWithJSON(&event); e != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "failed to bind",
		})
		return
	}

	eventPath, err := pr.CEs.CreateEventService(c.Request.Context(), event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: cuserror.InternalError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "event created success",
		Data:    eventPath,
	})
}
