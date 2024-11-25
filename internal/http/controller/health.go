package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	HealthController interface {
		Health(context *gin.Context)
	}

	healthController struct {
	}
)

func NewHealthController() HealthController {
	return &healthController{}
}

func (controller *healthController) Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
