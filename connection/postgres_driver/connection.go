package postgres_driver

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/models"
	"github.com/urcop/emotionalTracker/domain/repositories"
)

type connection struct {
	db *gorm.DB

	userRepository repositories.User
}

func makeConnection(db *gorm.DB) *connection {
	return &connection{
		db:             db,
		userRepository: &userRepository{db},
	}
}

func Make(user, password, host, port, database, sslmode string) (domain.Connection, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		database,
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to open database due [%w]", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to get DB object due [%w]", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping DB due [%w]", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("unable to migrate user table due [%w]", err)
	}

	return makeConnection(db), nil
}

func (c connection) User() repositories.User {
	return c.userRepository
}
