package controller

import (
	"net/http"

	"github.com/dshestapalau/gogotask/register/internal/domain/model"
	"github.com/dshestapalau/gogotask/register/internal/domain/service"
	"github.com/dshestapalau/gogotask/register/internal/http/message/request"
	"github.com/dshestapalau/gogotask/register/internal/http/message/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	UserController interface {
		Create(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) Create(ctx *gin.Context) {
	var request request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {

		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	controller.userService.CreateUser(model.UserCredentials{
		Login:    request.Login,
		Password: request.Password,
	})

	response := response.CreateUserResponse{
		Token: uuid.New().String(),
	}

	ctx.JSON(http.StatusOK, response)
}