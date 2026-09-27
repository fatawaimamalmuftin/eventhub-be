package handler

import (
	"context"
	"errors"
	"net/http"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/gin-gonic/gin"
)

type ILoginSrv interface {
	LoginSrv(c context.Context, user *dto.Account) error
}

type LoginHdr struct {
	Lh ILoginSrv
}

func LoginHandler(ls ILoginSrv) *LoginHdr {
	return &LoginHdr{
		Lh: ls,
	}
}

func (l *LoginHdr) Login(c *gin.Context) {
	logind := dto.Account{}

	if e := c.ShouldBindBodyWithJSON(&logind); e != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: cuserror.ErrBinding,
		})
		return
	}

	if err := l.Lh.LoginSrv(c.Request.Context(), &logind); err != nil {
		if err == cuserror.LenPas {
			c.JSON(http.StatusBadRequest, dto.Res{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, cuserror.InvalidEmail) {
			c.JSON(http.StatusBadRequest, dto.Res{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, cuserror.InvalidCredential) {
			c.JSON(http.StatusUnauthorized, dto.Res{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		// if err == pgx.ErrNoRows {
		// 	c.JSON(http.StatusInternalServerError, dto.Res{
		// 		Status:  false,
		// 		Message: err.Error(),
		// 	})
		// 	return
		// }

		if errors.Is(err, cuserror.InvalidHash) {
			c.JSON(http.StatusInternalServerError, dto.Res{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, cuserror.InternalError) {
			c.JSON(http.StatusInternalServerError, dto.Res{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: cuserror.InternalError.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "login success",
	})
}
