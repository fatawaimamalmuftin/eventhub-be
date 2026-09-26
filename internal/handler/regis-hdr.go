package handler

import (
	"context"
	"net/http"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/gin-gonic/gin"
)

type IRegisSrv interface {
	RegisSrv(newUser *dto.Regis, c context.Context) error
}

type RegisHdr struct {
	Rh IRegisSrv
}

func RegisHandler(rs IRegisSrv) *RegisHdr {
	return &RegisHdr{
		Rh: rs,
	}
}

func (r *RegisHdr) Regis(c *gin.Context) {
	newUser := dto.Regis{}

	if e := c.ShouldBindBodyWithJSON(&newUser); e != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "failed to bind",
		})
		return
	}

	if e := r.Rh.RegisSrv(&newUser, c.Request.Context()); e != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: e.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "registration successful",
		Data:    newUser.FullName,
	})
}
