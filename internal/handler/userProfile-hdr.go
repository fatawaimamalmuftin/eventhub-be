package handler

import (
	"context"
	"log"
	"net/http"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	jwtpkg "github.com/fatawaimamalmuftin/eventhub-be/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type IUserProfile interface {
	GetUserProfileSrv(c context.Context, userId int) (*model.UserProfile, error)
}

type UserProfileHdr struct {
	UPs IUserProfile
}

func UserProfileHandler(ups IUserProfile) *UserProfileHdr {
	return &UserProfileHdr{
		UPs: ups,
	}
}

func (u *UserProfileHdr) GetUserProfileHandler(c *gin.Context) {

	tokenClaims, exists := c.Get("tokenCleims")

	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "please login first",
		})
		return
	}

	userId, ok := tokenClaims.(jwtpkg.JWTclem)

	if !ok {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid token",
		})
		return
	}

	profile, err := u.UPs.GetUserProfileSrv(c.Request.Context(), userId.Id)

	if err != nil {
		log.Println(err.Error())

		if err == cuserror.UserNotFound {
			c.JSON(http.StatusNotFound, dto.Res{
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
		Message: "success get user profile",
		Data:    profile,
	})
}
