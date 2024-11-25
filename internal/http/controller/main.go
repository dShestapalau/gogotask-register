package controller

type Controller struct {
	HealthController HealthController
}

func NewController() *Controller {
	return &Controller{
		HealthController: NewHealthController(),
	}
}
