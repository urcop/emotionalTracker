package update_user

import (
	"errors"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/models"
)

type Request struct {
	User *models.User `json:"user"`
}

type Response struct {
	Success bool `json:"success"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	if req.User == nil {
		return nil, errors.New("user is required")
	}

	if req.User.Id == "" {
		return nil, errors.New("user id is required")
	}

	err := c.Connection().User().Update(req.User)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return &Response{Success: true}, nil
}
