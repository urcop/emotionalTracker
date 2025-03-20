package get_horoscope

import (
	"errors"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/services/zodiac"
)

type Request struct {
	ZodiacSign string `json:"zodiac_sign"`
	Birthdate  string `json:"birthdate"`
}

type Response struct {
	Date          string `json:"date"`
	HoroscopeData string `json:"horoscope_data"`
	ZodiacSign    string `json:"zodiac_sign"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	var sign string

	// If zodiac sign is provided directly, use it
	if req.ZodiacSign != "" {
		sign = req.ZodiacSign
	} else if req.Birthdate != "" {
		// If birthdate is provided, derive the zodiac sign
		zodiacSign, err := zodiac.ParseZodiacSign(req.Birthdate)
		if err != nil {
			return nil, errors.New("failed to parse zodiac sign from birthdate: " + err.Error())
		}
		sign = zodiacSign.Name
	} else {
		return nil, errors.New("either zodiac_sign or birthdate must be provided")
	}

	// Get the horoscope from the service
	horoscope, err := c.Services().Horoscope().GetDailyHoroscope(sign)
	if err != nil {
		return nil, errors.New("failed to get horoscope: " + err.Error())
	}

	return &Response{
		Date:          horoscope.Date,
		HoroscopeData: horoscope.HoroscopeData,
		ZodiacSign:    horoscope.ZodiacSign,
	}, nil
}
