package app

import (
	"github.com/dshestapalau/gogotask/register/internal/http/controller"
	"github.com/dshestapalau/gogotask/register/internal/http/router"
	utils "github.com/dshestapalau/gogotask/register/pkg/http"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	engine := gin.Default()

	utils.RegisterValidationRules()

	controller := controller.NewController()
	router.New(engine, controller)

	engine.Run(":9000")
}
