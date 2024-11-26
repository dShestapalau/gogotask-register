package controller

import "github.com/dshestapalau/gogotask/register/internal/domain/service"

type Controller struct {
	HealthController HealthController
	UserController   UserController
}

func NewController(userService service.UserService) *Controller {
	return &Controller{
		HealthController: NewHealthController(),
		UserController:   NewUserController(userService),
	}
}
