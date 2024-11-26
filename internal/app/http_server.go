package app

import (
	"fmt"

	"github.com/dshestapalau/gogotask/register/internal/config"
	"github.com/dshestapalau/gogotask/register/internal/http/controller"
	"github.com/dshestapalau/gogotask/register/internal/http/router"
	utils "github.com/dshestapalau/gogotask/register/pkg/http"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	configuration := config.NewConfiguration()
	configuration.LoadConfig()

	fmt.Println(configuration)

	dbConnection := config.OpenConnection(configuration)
	defer config.CloseDatabaseConnection(dbConnection)

	engine := gin.Default()

	utils.RegisterValidationRules()

	controller := controller.NewController()
	router.New(engine, controller)

	engine.Run(":" + configuration.ServerPort)
}
