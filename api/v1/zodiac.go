package v1

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/urcop/emotionalTracker/domain"
	"github.com/urcop/emotionalTracker/domain/cases/get_horoscope"
	"github.com/urcop/emotionalTracker/domain/cases/get_zodiac_sign"
)

// GetZodiacSign godoc
// @Summary Get zodiac sign by birthdate
// @Description Get zodiac sign information by birthdate
// @Tags zodiac
// @Accept json
// @Produce json
// @Param birthdate query string true "Birthdate in format DD.MM.YYYY or YYYY-MM-DD"
// @Success 200 {object} RawResponse "Returns zodiac sign information"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/zodiac/ [get]
func GetZodiacSign(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	birthdate := ctx.Query("birthdate")

	if birthdate == "" {
		return BadRequest(nil)
	}

	var req get_zodiac_sign.Request
	req.Birthdate = birthdate

	resp, err := get_zodiac_sign.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}

// GetUserZodiacSign godoc
// @Summary Get user's zodiac sign
// @Description Get zodiac sign information for a specific user
// @Tags zodiac
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} RawResponse "Returns zodiac sign information"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 404 {object} RawResponse "User not found"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/{telegram_id}/zodiac [get]
func GetUserZodiacSign(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	userId := ctx.Params("telegram_id")

	if userId == "" {
		return BadRequest(nil)
	}

	user, err := c.Connection().User().GetUserByTelegramId(userId)
	if err != nil {
		return BadRequest(errors.New("user not found"))
	}

	if user.Birthday == nil || *user.Birthday == "" {
		return BadRequest(errors.New("user doesn't have a birthday set"))
	}

	var req get_zodiac_sign.Request
	req.Birthdate = *user.Birthday

	resp, err := get_zodiac_sign.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}

// GetHoroscope godoc
// @Summary Get daily horoscope for a zodiac sign or by birthdate
// @Description Get translated daily horoscope for a zodiac sign or by birthdate
// @Tags zodiac
// @Accept json
// @Produce json
// @Param sign query string false "Zodiac sign name (e.g., Aries, Taurus, etc.)"
// @Param birthdate query string false "Birthdate in format DD.MM.YYYY or YYYY-MM-DD"
// @Success 200 {object} RawResponse "Returns translated horoscope information"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/zodiac/horoscope/ [get]
func GetHoroscope(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	sign := ctx.Query("sign")
	birthdate := ctx.Query("birthdate")

	if sign == "" && birthdate == "" {
		return BadRequest(errors.New("either sign or birthdate must be provided"))
	}

	var req get_horoscope.Request
	req.ZodiacSign = sign
	req.Birthdate = birthdate

	resp, err := get_horoscope.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}

// GetUserHoroscope godoc
// @Summary Get daily horoscope for a user
// @Description Get translated daily horoscope for a user based on their birthdate
// @Tags zodiac
// @Accept json
// @Produce json
// @Param id path string true "User ID or Telegram ID"
// @Success 200 {object} RawResponse "Returns translated horoscope information"
// @Failure 400 {object} RawResponse "Bad request"
// @Failure 404 {object} RawResponse "User not found"
// @Failure 500 {object} RawResponse "Internal server error"
// @Router /api/v1/user/{telegram_id}/horoscope [get]
func GetUserHoroscope(c domain.Context, ctx *fiber.Ctx) *RawResponse {
	userId := ctx.Params("telegram_id")

	if userId == "" {
		return BadRequest(nil)
	}

	user, err := c.Connection().User().GetUserByTelegramId(userId)
	if err != nil {
		return BadRequest(errors.New("user not found"))
	}

	if user.Birthday == nil || *user.Birthday == "" {
		return BadRequest(errors.New("user doesn't have a birthday set"))
	}

	var req get_horoscope.Request
	req.Birthdate = *user.Birthday

	resp, err := get_horoscope.Run(c, req)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(resp)
}
