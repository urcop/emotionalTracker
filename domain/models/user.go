package models

import "time"

type User struct {
	Id         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	TelegramId string    `json:"telegram_id"`
	Username   *string   `json:"username"`
	Birthday   *string   `json:"birthsday"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
