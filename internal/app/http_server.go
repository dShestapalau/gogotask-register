package app

import (
	"github.com/dshestapalau/gogotask/register/internal/config"
	"github.com/dshestapalau/gogotask/register/internal/domain/service"
	"github.com/dshestapalau/gogotask/register/internal/http/controller"
	"github.com/dshestapalau/gogotask/register/internal/http/router"
	"github.com/dshestapalau/gogotask/register/internal/persistence/repository"
	utils "github.com/dshestapalau/gogotask/register/pkg/http"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	configuration := config.NewConfiguration()
	configuration.LoadConfig()

	dbConnection := config.OpenConnection(configuration)
	defer config.CloseDatabaseConnection(dbConnection)

	engine := gin.Default()

	utils.RegisterValidationRules()

	userCredentialsRepository := repository.NewUserCredentialsRepository(dbConnection)
	userService := service.NewUserService(userCredentialsRepository)

	controller := controller.NewController(userService)
	router.New(engine, controller)

	engine.Run(":" + configuration.ServerPort)
}
