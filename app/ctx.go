package app

import (
	"log"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/services"
	"github.com/urcop/emotionalTracker/services/config"
	"github.com/urcop/emotionalTracker/services/horoscope"
	"github.com/urcop/emotionalTracker/services/logger"
)

type ctx struct {
	services   domain.Services
	connection domain.Connection
}

type svs struct {
	config    services.Config
	logger    services.Logger
	horoscope services.Horoscope
}

func (s *svs) Logger() services.Logger {
	return s.logger
}

func (s *svs) Config() services.Config {
	return s.config
}

func (s *svs) Horoscope() services.Horoscope {
	return s.horoscope
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

	horoscopeService, err := horoscope.NewHoroscopeService()
	if err != nil {
		log.Fatalf("cant initialize horoscope service due [%s]", err)
	}

	return &ctx{
		services: &svs{
			config:    cfg,
			logger:    logger.Init(cfg.EnvLevel()),
			horoscope: horoscopeService,
		},
		connection: connection,
	}
}
