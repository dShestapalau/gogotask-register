package router

import (
	"github.com/dshestapalau/gogotask/register/internal/http/controller"
	"github.com/gin-gonic/gin"
)

func New(engine *gin.Engine, controllers *controller.Controller) {
	routes := engine.Group("/api/v1/user")

	{
		routes.GET("/health", controllers.HealthController.Health)
	}

}
