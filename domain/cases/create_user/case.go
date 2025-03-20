package create_user

import (
	"errors"

	"github.com/google/uuid"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/models"
)

type Request struct {
	User *models.User `json:"user"`
}

type Response struct {
	Id string `json:"id"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	id := uuid.New().String()

	user := models.User{
		Id:         id,
		FirstName:  req.User.FirstName,
		SecondName: req.User.SecondName,
		Username:   req.User.Username,
		Birthday:   req.User.Birthday,
		TelegramId: req.User.TelegramId,
	}

	id, err := c.Connection().User().Insert(&user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return &Response{Id: id}, nil
}
