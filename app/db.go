package app

import (
	"github.com/urcop/emotionalTracker/connection/postgres_driver"
	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/services/config"
)

func InitDB(cfg *config.Config) (domain.Connection, error) {
	return postgres_driver.Make(cfg.PostgresUser(), cfg.PostgresPassword(), cfg.PostgresHost(), cfg.PostgresPort(), cfg.PostgresName(), cfg.SslMode())
}
