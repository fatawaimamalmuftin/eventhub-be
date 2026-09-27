package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	"github.com/gin-gonic/gin"
)

type IEventsDetailSrv interface {
	GetEventDetail(
		id int,
		c context.Context,
	) (model.EventDetail, error)
}

type EventDetail struct {
	EDh IEventsDetailSrv
}

func EventDetailHandler(edh IEventsDetailSrv) *EventDetail {
	return &EventDetail{
		EDh: edh,
	}
}

func (e *EventDetail) GetEventDetail(c *gin.Context) {

	eventID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "invalid event id",
		})
		return
	}

	event, err := e.EDh.GetEventDetail(
		eventID,
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "event detail retrieved successfully",
		Data:    event,
	})
}
