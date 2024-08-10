package handlers

import "github.com/FoodMoodOTG/examplearch/domain"

type Handler struct {
	//sso_v1.UnimplementedExampleServiceServer
	context domain.Context
}

func NewHandler(context domain.Context) *Handler {
	return &Handler{context: context}
}
