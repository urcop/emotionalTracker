package v1

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/urcop/emotionalTracker/domain"
)

type RawResponse struct {
	error      error
	status     int
	additional interface{}
	payload    interface{}
}

func (r *RawResponse) Error() error {
	return r.error
}

func (r *RawResponse) WithPayload(payload any) *RawResponse {
	r.payload = payload
	return r
}

func (r *RawResponse) Body() *ResponseBody {
	return &ResponseBody{
		Response: Response{
			Status: r.status,
		},
		Additional: r.additional,
		Payload:    r.payload,
	}
}

type Response struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message,omitempty" example:"some message"`
}

type ResponseBody struct {
	Response   `json:"response"`
	Additional interface{} `json:"-"`
	Payload    interface{} `json:"payload,omitempty"`
}

func WrapHandler(handler func(c domain.Context, ctx *fiber.Ctx) *RawResponse) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newContext, ok := ctx.Locals("context").(domain.Context)
		if !ok {
			return nil
		}

		response := handler(newContext, ctx)
		body := response.Body()

		status := body.Status

		if err := response.Error(); err != nil {

			body.Message = response.Error().Error()
		}
		return ctx.Status(status).JSON(body)
	}
}

func WrapHandlerOAuth(handler func(c domain.Context, ctx *fiber.Ctx) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		domainCtx, ok := ctx.Locals("context").(domain.Context)
		if !ok {
			return nil
		}

		err := handler(domainCtx, ctx)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return nil
	}
}

func BadRequest(err error) *RawResponse {
	return &RawResponse{
		status: http.StatusBadRequest,
		error:  err,
	}
}

func InternalServerError(err error) *RawResponse {
	return &RawResponse{
		status: http.StatusInternalServerError,
		error:  err,
	}
}

func OK(payload any) *RawResponse {
	out := &RawResponse{
		status: http.StatusOK,
	}

	if payload != nil {
		out.WithPayload(payload)
	}

	return out
}
