package domain

import "github.com/FoodMoodOTG/examplearch/domain/services"

type Services interface {
	Config() services.Config
}
