package get_zodiac_sign

import (
	"errors"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/services/zodiac"
)

type Request struct {
	Birthdate string `json:"birthdate"`
}

type Response struct {
	ZodiacSign *zodiac.ZodiacSign `json:"zodiac_sign"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	if req.Birthdate == "" {
		return nil, errors.New("birthdate is required")
	}

	sign, err := zodiac.ParseZodiacSign(req.Birthdate)
	if err != nil {
		return nil, errors.New("failed to parse zodiac sign: " + err.Error())
	}

	return &Response{ZodiacSign: sign}, nil
}
