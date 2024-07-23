package server

import (
	"github.com/FoodMoodOTG/examplearch/domain"
	"github.com/FoodMoodOTG/examplearch/domain/services"
)

type ctx struct {
	services   domain.Services
	connection domain.Connection
}

type svs struct {
	config services.Config
}

func (s *svs) Config() services.Config {
	return s.config
}

func (c *ctx) Services() domain.Services {
	return c.services
}

func (c *ctx) Connection() domain.Connection {
	return c.connection
}

func (c *ctx) Make() domain.Context {
	return &ctx{
		services:   c.services,
		connection: c.connection,
	}
}
