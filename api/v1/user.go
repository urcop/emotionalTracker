package v1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/cases/create_user"
	"github.com/urcop/emotionalTracker/domain/cases/delete_user"
	"github.com/urcop/emotionalTracker/domain/cases/get_all_users"
	"github.com/urcop/emotionalTracker/domain/cases/get_user"
	"github.com/urcop/emotionalTracker/domain/cases/get_user_by_telegram_id"
	"github.com/urcop/emotionalTracker/domain/cases/update_user"
	"github.com/urcop/emotionalTracker/domain/cases/update_user_birthday"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept  json
// @Produce  json
// @Param  create_user.Request  body  create_user.Request  true  "User request"
// @Success 200 {object} RawResponse "Returns the ID of the created user"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/ [post]
func CreateUser(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	var req create_user.Request

	if err := ctx.BodyParser(&req); err != nil {
		return BadRequest(err)
	}

	id, err := create_user.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(id)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce  json
// @Success 200 {object} RawResponse "Returns a list of users"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/ [get]
func GetAllUsers(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	var req get_all_users.Request

	users, err := get_all_users.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(users)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Retrieve a single user by its ID
// @Tags users
// @Produce  json
// @Param  id  path  string  true  "User ID"
// @Success 200 {object} RawResponse "Returns the requested user"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/{id}/ [get]
func GetUser(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	id := ctx.Params("id")

	var req get_user.Request
	req.Id = id

	user, err := get_user.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user with the provided details
// @Tags users
// @Accept  json
// @Produce  json
// @Param  update_user.Request  body  update_user.Request  true  "User update request"
// @Success 200 {object} RawResponse "Returns success status"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/ [put]
func UpdateUser(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	var req update_user.Request

	if err := ctx.BodyParser(&req); err != nil {
		return BadRequest(err)
	}

	resp, err := update_user.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by its ID
// @Tags users
// @Produce  json
// @Param  id  path  string  true  "User ID"
// @Success 200 {object} RawResponse "Returns success status"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/{id}/ [delete]
func DeleteUser(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	id := ctx.Params("id")

	var req delete_user.Request
	req.Id = id

	resp, err := delete_user.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}

// GetUserByTelegramId godoc
// @Summary Get a user by Telegram ID
// @Description Retrieve a single user by its Telegram ID
// @Tags users
// @Produce  json
// @Param  telegram_id  query  string  true  "Telegram ID"
// @Success 200 {object} RawResponse "Returns the requested user"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/telegram/ [get]
func GetUserByTelegramId(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	telegramId := ctx.Query("telegram_id")

	if telegramId == "" {
		return BadRequest(nil)
	}

	var req get_user_by_telegram_id.Request
	req.TelegramId = telegramId

	resp, err := get_user_by_telegram_id.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}

// UpdateUserBirthday godoc
// @Summary Update user's birthdate
// @Description Update the birthdate for a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param birthdate body update_user_birthday.Request true "Birthdate information"
// @Success 200 {object} RawResponse "Returns success status"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/{id}/birthday/ [put]
func UpdateUserBirthday(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	userId := ctx.Params("id")

	if userId == "" {
		return BadRequest(nil)
	}

	var req update_user_birthday.Request

	if err := ctx.BodyParser(&req); err != nil {
		return BadRequest(err)
	}

	req.UserId = userId

	resp, err := update_user_birthday.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}
