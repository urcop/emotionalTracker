package delete_user

import (
	"errors"

	"github.com/urcop/emotionalTracker/domain"
)

type Request struct {
	Id string `json:"id"`
}

type Response struct {
	Success bool `json:"success"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	if req.Id == "" {
		return nil, errors.New("id is required")
	}

	// Сначала получаем пользователя, чтобы убедиться, что он существует
	user, err := c.Connection().User().GetUser(req.Id)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	err = c.Connection().User().Delete(user)
	if err != nil {
		return nil, errors.New("failed to delete user")
	}

	return &Response{Success: true}, nil
}
