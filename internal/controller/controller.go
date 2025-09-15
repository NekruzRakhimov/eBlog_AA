package controller

import (
	"eBlog/internal/contracts"
	"eBlog/internal/service"
)

type Controller struct {
	service contracts.ServiceI
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		service: svc,
	}
}
