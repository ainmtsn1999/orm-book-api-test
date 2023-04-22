package controller

import "github.com/ainmtsn1999/orm-book-api-test/service"

type Controller struct {
	service service.ServiceInterface
}

func NewController(service service.ServiceInterface) *Controller {
	return &Controller{service: service}
}
