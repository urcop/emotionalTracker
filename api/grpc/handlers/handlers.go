package handlers

import "github.com/urcop/emotionalTracker/domain"

type Handler struct {
	//sso_v1.UnimplementedExampleServiceServer
	context domain.Context
}

func NewHandler(context domain.Context) *Handler {
	return &Handler{context: context}
}
