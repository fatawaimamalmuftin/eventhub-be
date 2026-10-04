package handler

import (
	"context"
	"net/http"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/model"
	jwtpkg "github.com/fatawaimamalmuftin/eventhub-be/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type ITestimonialSrv interface {
	GetTestimonials(c context.Context) ([]model.Testimonial, error)
	CreateTestimonial(c context.Context, data dto.CreateTestimonial, userID int) error
}

type TestimonialHandler struct {
	TS ITestimonialSrv
}

func ProviderTestimonialHandler(ts ITestimonialSrv) *TestimonialHandler {
	return &TestimonialHandler{
		TS: ts,
	}
}

func (h *TestimonialHandler) GetTestimonials(c *gin.Context) {
	testimonials, err := h.TS.GetTestimonials(
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
		Message: "testimonials retrieved successfully",
		Data:    testimonials,
	})
}

func (h *TestimonialHandler) CreateTestimonial(c *gin.Context) {
	var data dto.CreateTestimonial

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "invalid request body",
		})
		return
	}

	tokenClaims, exists := c.Get("tokenCleims")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "unauthorized",
		})
		return
	}

	claims, ok := tokenClaims.(jwtpkg.JWTclem)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid token claims",
		})
		return
	}

	err := h.TS.CreateTestimonial(c.Request.Context(), data, claims.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "failed to create testimonial",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Res{
		Status:  true,
		Message: "testimonial created successfully",
	})
}
