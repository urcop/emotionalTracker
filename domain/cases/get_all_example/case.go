package get_all_example

import (
	"fmt"
	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/models"
)

type Response struct {
	Examples []*models.Example `json:"examples"`
}

func Run(c domain.Context) (*Response, error) {
	examples, err := c.Connection().Example().All()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve roles due [%s]", err)
	}

	return &Response{Examples: examples}, nil
}
