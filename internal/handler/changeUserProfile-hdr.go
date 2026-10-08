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

// ChangeUserProfileHdr godoc
// @Summary      Change user profile
// @Description  Update user profile details for the authenticated user
// @Tags         Change User Profile
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Param        request  body      dto.ChangeUserProfile  true  "Change profile payload"
// @Success      200      {object}  dto.Res{data=string}   "profile updated successfully"
// @Failure      400      {object}  dto.Res                "bad request error message"
// @Failure      401      {object}  dto.Res                "invalid token"
// @Failure      500      {object}  dto.Res                "internal server error"
// @Router       /events/changeuserprofile [patch]
func (c *ChangeUserProfileHdrS) ChangeUserProfileHdr(ctx *gin.Context) {
	var userProfile dto.ChangeUserProfile

	if err := ctx.ShouldBindJSON(&userProfile); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: err.Error(),
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
