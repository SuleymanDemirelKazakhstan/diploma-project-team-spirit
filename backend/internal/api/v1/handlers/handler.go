package handlers

import "secondChance/internal/services"

type Handler struct {
	ServiceLayer *services.Layer
}

func NewHandler(serviceLayer *services.Layer) *Handler {
	return &Handler{
		ServiceLayer: serviceLayer,
	}
}
