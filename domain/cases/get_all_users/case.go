package get_all_users

import (
	"errors"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/models"
)

type Request struct {
	// Пустой запрос, так как нет параметров
}

type Response struct {
	Users []*models.User `json:"users"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	users, err := c.Connection().User().All()
	if err != nil {
		return nil, errors.New("failed to get users")
	}

	return &Response{Users: users}, nil
}
