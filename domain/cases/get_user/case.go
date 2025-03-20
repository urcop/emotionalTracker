package get_user

import (
	"errors"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/models"
)

type Request struct {
	Id string `json:"id"`
}

type Response struct {
	User *models.User `json:"user"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	if req.Id == "" {
		return nil, errors.New("id is required")
	}

	user, err := c.Connection().User().GetUser(req.Id)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	return &Response{User: user}, nil
}
