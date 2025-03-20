package domain

import "github.com/urcop/emotionalTracker/domain/services"

type Services interface {
	Config() services.Config
	Logger() services.Logger
	Horoscope() services.Horoscope
}
