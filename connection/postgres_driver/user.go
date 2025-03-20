package postgres_driver

import (
	"time"

	"gorm.io/gorm"

	"github.com/urcop/emotionalTracker/domain/models"
)

type userRepository struct {
	db *gorm.DB
}

type user struct {
	Id         string
	FirstName  string
	SecondName string
	TelegramId string
	Username   *string
	Birthday   *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u user) model() *models.User {
	return &models.User{
		Id:         u.Id,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		TelegramId: u.TelegramId,
		Username:   u.Username,
		Birthday:   u.Birthday,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func makeUser(u *models.User) user {
	return user{
		Id:         u.Id,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		TelegramId: u.TelegramId,
		Username:   u.Username,
		Birthday:   u.Birthday,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (u *userRepository) Insert(example *models.User) (string, error) {
	userModel := makeUser(example)

	if err := u.db.Create(&userModel).Error; err != nil {
		return "", err
	}

	return userModel.Id, nil
}

func (u *userRepository) GetUser(id string) (*models.User, error) {
	var result user

	if err := u.db.Where(user{Id: id}).First(&result).Error; err != nil {
		return nil, err
	}

	return result.model(), nil
}

func (u *userRepository) All() ([]*models.User, error) {
	var result []user

	if err := u.db.Find(&result).Error; err != nil {
		return nil, err
	}

	out := make([]*models.User, len(result))

	for i, em := range result {
		out[i] = em.model()
	}

	return out, nil
}

func (u *userRepository) Update(userModel *models.User) error {
	dbUser := makeUser(userModel)
	return u.db.Model(&user{}).Where(user{Id: userModel.Id}).Updates(dbUser).Error
}

func (u *userRepository) Delete(user *models.User) error {
	userModel := makeUser(user)
	return u.db.Delete(&userModel).Error
}

func (u *userRepository) GetUserByTelegramId(telegramId string) (*models.User, error) {
	var result user

	if err := u.db.Where(user{TelegramId: telegramId}).First(&result).Error; err != nil {
		return nil, err
	}

	return result.model(), nil
}
