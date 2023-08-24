package handler

import "github.com/zpix1/avito-test-task/pkg/service"

type Handler struct {
	service service.Implementation
}

func NewHandler(service service.Implementation) *Handler {
	return &Handler{
		service: service,
	}
}
