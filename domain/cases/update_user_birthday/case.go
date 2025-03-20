package update_user_birthday

import (
	"errors"
	"time"

	"github.com/urcop/emotionalTracker/domain"
)

type Request struct {
	UserId    string `json:"user_id"`
	Birthdate string `json:"birthdate"`
}

type Response struct {
	Success bool `json:"success"`
}

func Run(c domain.Context, req Request) (*Response, error) {
	if req.UserId == "" {
		return nil, errors.New("user id is required")
	}

	if req.Birthdate == "" {
		return nil, errors.New("birthdate is required")
	}

	// Проверим, что дата рождения в правильном формате
	_, err := validateBirthdate(req.Birthdate)
	if err != nil {
		return nil, err
	}

	// Получаем пользователя из БД
	user, err := c.Connection().User().GetUser(req.UserId)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	// Обновляем дату рождения
	user.Birthday = &req.Birthdate

	// Сохраняем пользователя
	err = c.Connection().User().Update(user)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return &Response{Success: true}, nil
}

func validateBirthdate(birthdate string) (time.Time, error) {
	var parsedDate time.Time
	var err error

	if len(birthdate) == 10 && birthdate[2] == '.' && birthdate[5] == '.' {
		parsedDate, err = time.Parse("02.01.2006", birthdate)
		if err != nil {
			return time.Time{}, errors.New("invalid date format, expected DD.MM.YYYY")
		}
	} else if len(birthdate) == 10 && birthdate[4] == '-' && birthdate[7] == '-' {
		parsedDate, err = time.Parse("2006-01-02", birthdate)
		if err != nil {
			return time.Time{}, errors.New("invalid date format, expected YYYY-MM-DD")
		}
	} else {
		return time.Time{}, errors.New("invalid date format, expected DD.MM.YYYY or YYYY-MM-DD")
	}

	if parsedDate.After(time.Now()) {
		return time.Time{}, errors.New("birthdate cannot be in the future")
	}

	maxAge := time.Now().AddDate(-120, 0, 0)
	if parsedDate.Before(maxAge) {
		return time.Time{}, errors.New("birthdate is too old (more than 120 years)")
	}

	return parsedDate, nil
}
