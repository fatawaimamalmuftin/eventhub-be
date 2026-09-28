package handler

import (
	"net/http"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	jwtpkg "github.com/fatawaimamalmuftin/eventhub-be/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type IChangeUserProfile interface {
	ChangeUserProfileSrv(userID int, data dto.ChangeUserProfile) (string, error)
}

type ChangeUserProfileHdrS struct {
	CUPs IChangeUserProfile
}

func ChangeUserProfileHandler(cups IChangeUserProfile) *ChangeUserProfileHdrS {
	return &ChangeUserProfileHdrS{
		CUPs: cups,
	}
}

func (c *ChangeUserProfileHdrS) ChangeUserProfileHdr(ctx *gin.Context) {
	var userProfile dto.ChangeUserProfile

	if err := ctx.ShouldBindJSON(&userProfile); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "failed to bind",
		})
		return
	}

	tokenClaims, exists := ctx.Get("tokenCleims")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid token",
		})
		return
	}

	claims, ok := tokenClaims.(jwtpkg.JWTclem)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "internal server error",
		})
		return
	}

	profilePath, err := c.CUPs.ChangeUserProfileSrv(claims.Id, userProfile)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "profile updated successfully",
		Data:    profilePath,
	})
}
