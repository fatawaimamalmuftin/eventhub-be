package handler

import (
	"context"
	"net/http"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type IRegisSrv interface {
	RegisSrv(newUser *model.Regis, c context.Context) error
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
	newUser := model.Regis{}

	if e := c.ShouldBindBodyWithJSON(&newUser); e != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "failed to bind",
		})
		return
	}

	if e := r.Rh.RegisSrv(&newUser, c.Request.Context()); e != nil {
		if e == cuserror.LenPas {
			c.JSON(http.StatusBadRequest, dto.Res{
				Status:  false,
				Message: e.Error(),
			})
			return
		}

		if e == cuserror.ErrNoRowAffected {
			c.JSON(http.StatusInternalServerError, dto.Res{
				Status:  false,
				Message: e.Error(),
			})
			return
		}

		if e == cuserror.ErrIsExist {
			c.JSON(http.StatusConflict, dto.Res{
				Status:  false,
				Message: e.Error(),
			})
			return
		}

		if e == pgx.ErrNoRows {
			c.JSON(http.StatusInternalServerError, dto.Res{
				Status:  false,
				Message: e.Error(),
			})
		}
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "registration successful",
		Data:    newUser.FullName,
	})
}
