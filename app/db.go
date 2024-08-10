package app

import (
	"github.com/FoodMoodOTG/examplearch/connection/postgres_driver"
	"github.com/FoodMoodOTG/examplearch/domain"
	"github.com/FoodMoodOTG/examplearch/services/config"
)

func InitDB(cfg *config.Config) (domain.Connection, error) {
	return postgres_driver.Make(cfg.PostgresUser(), cfg.PostgresPassword(), cfg.PostgresHost(), cfg.PostgresPort(), cfg.PostgresName(), cfg.SslMode())
}
