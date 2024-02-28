package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/matisiekpl/electrocardiogram-server/internal/service"
)

type Controllers interface {
	Info() InfoController

	Route(e *echo.Echo)
}

type controllers struct {
	infoController InfoController
}

func NewControllers(services service.Services) Controllers {
	infoController := newInfoController()
	return &controllers{
		infoController: infoController,
	}

}

func (c controllers) Info() InfoController {
	return c.infoController
}

func (c controllers) Route(e *echo.Echo) {
	e.GET("/", c.infoController.Info)
}
