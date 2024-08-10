package app

import (
	"github.com/FoodMoodOTG/examplearch/domain"
	"github.com/FoodMoodOTG/examplearch/domain/services"
	"github.com/FoodMoodOTG/examplearch/services/config"
	"github.com/FoodMoodOTG/examplearch/services/logger"
	"log"
)

type ctx struct {
	services   domain.Services
	connection domain.Connection
}

type svs struct {
	config services.Config
	logger services.Logger
}

func (s *svs) Logger() services.Logger {
	return s.logger
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

func InitCtx() *ctx {
	cfg := config.Make()
	connection, err := InitDB(cfg)
	if err != nil {
		log.Fatalf("cant initialize connection context due [%s]", err)
	}

	return &ctx{
		services: &svs{
			config: cfg,
			logger: logger.Init(cfg.EnvLevel()),
		},
		connection: connection,
	}
}
