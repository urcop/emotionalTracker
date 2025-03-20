package get_user_by_telegram_id

import (
	"errors"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/models"
)

type Request struct {
	TelegramId string `json:"telegram_id"`
}

type Response struct {
	User *models.User `json:"user"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	if req.TelegramId == "" {
		return nil, errors.New("telegram_id is required")
	}

	user, err := c.Connection().User().GetUserByTelegramId(req.TelegramId)
	if err != nil {
		return nil, errors.New("failed to get user by telegram id")
	}

	return &Response{User: user}, nil
}
