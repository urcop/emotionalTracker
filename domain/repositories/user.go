package repositories

import "github.com/urcop/emotionalTracker/domain/models"

type User interface {
	Insert(user *models.User) (string, error)
	GetUser(id string) (*models.User, error)
	All() ([]*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
	GetUserByTelegramId(telegramId string) (*models.User, error)
}
