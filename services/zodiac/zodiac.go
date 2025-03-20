package zodiac

import (
	"fmt"
	"time"
)

type ZodiacSign struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Element   string `json:"element"`
	Symbol    string `json:"symbol"`
}

var zodiacSigns = []ZodiacSign{
	{
		Name:      "Aries",
		StartDate: "21.03",
		EndDate:   "19.04",
		Element:   "Fire",
		Symbol:    "♈",
	},
	{
		Name:      "Taurus",
		StartDate: "20.04",
		EndDate:   "20.05",
		Element:   "Earth",
		Symbol:    "♉",
	},
	{
		Name:      "Gemini",
		StartDate: "21.05",
		EndDate:   "20.06",
		Element:   "Air",
		Symbol:    "♊",
	},
	{
		Name:      "Cancer",
		StartDate: "21.06",
		EndDate:   "22.07",
		Element:   "Water",
		Symbol:    "♋",
	},
	{
		Name:      "Leo",
		StartDate: "23.07",
		EndDate:   "22.08",
		Element:   "Fire",
		Symbol:    "♌",
	},
	{
		Name:      "Virgo",
		StartDate: "23.08",
		EndDate:   "22.09",
		Element:   "Earth",
		Symbol:    "♍",
	},
	{
		Name:      "Libra",
		StartDate: "23.09",
		EndDate:   "22.10",
		Element:   "Air",
		Symbol:    "♎",
	},
	{
		Name:      "Scorpio",
		StartDate: "23.10",
		EndDate:   "21.11",
		Element:   "Water",
		Symbol:    "♏",
	},
	{
		Name:      "Sagittarius",
		StartDate: "22.11",
		EndDate:   "21.12",
		Element:   "Fire",
		Symbol:    "♐",
	},
	{
		Name:      "Capricorn",
		StartDate: "22.12",
		EndDate:   "19.01",
		Element:   "Earth",
		Symbol:    "♑",
	},
	{
		Name:      "Aquarius",
		StartDate: "20.01",
		EndDate:   "18.02",
		Element:   "Air",
		Symbol:    "♒",
	},
	{
		Name:      "Pisces",
		StartDate: "19.02",
		EndDate:   "20.03",
		Element:   "Water",
		Symbol:    "♓",
	},
}

func ParseZodiacSign(birthdate string) (*ZodiacSign, error) {
	var parsedDate time.Time
	var err error

	if len(birthdate) == 10 && birthdate[2] == '.' && birthdate[5] == '.' {
		parsedDate, err = time.Parse("02.01.2006", birthdate)
		if err != nil {
			return nil, fmt.Errorf("неверный формат даты, ожидается DD.MM.YYYY: %w", err)
		}
	} else if len(birthdate) == 10 && birthdate[4] == '-' && birthdate[7] == '-' {
		parsedDate, err = time.Parse("2006-01-02", birthdate)
		if err != nil {
			return nil, fmt.Errorf("неверный формат даты, ожидается YYYY-MM-DD: %w", err)
		}
	} else {
		return nil, fmt.Errorf("неверный формат даты, ожидается DD.MM.YYYY или YYYY-MM-DD")
	}

	month := parsedDate.Month()
	day := parsedDate.Day()

	for _, sign := range zodiacSigns {
		startMonth, startDay, err := parseDate(sign.StartDate)
		if err != nil {
			return nil, err
		}

		endMonth, endDay, err := parseDate(sign.EndDate)
		if err != nil {
			return nil, err
		}

		if startMonth == 12 && endMonth == 1 {
			if (month == 12 && day >= startDay) || (month == 1 && day <= endDay) {
				return &sign, nil
			}
		} else {
			if (month == startMonth && day >= startDay) || (month == endMonth && day <= endDay) {
				return &sign, nil
			}
		}
	}

	return nil, fmt.Errorf("не удалось определить знак зодиака для даты %s", birthdate)
}

func parseDate(date string) (month time.Month, day int, err error) {
	parsedDate, err := time.Parse("02.01", date)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка парсинга даты %s: %w", date, err)
	}
	return parsedDate.Month(), parsedDate.Day(), nil
}
