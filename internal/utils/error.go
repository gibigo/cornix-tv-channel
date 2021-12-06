package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewHTTPError(ctx *fiber.Ctx, status int, err interface{}) error {
	var er *HTTPError
	switch err := err.(type) {
	case error:
		er = &HTTPError{
			Code:    status,
			Message: err.Error(),
		}
	case string:
		er = &HTTPError{
			Code:    status,
			Message: err,
		}
	default:
		er = &HTTPError{
			Code:    status,
			Message: fmt.Sprintf("%s", err),
		}
	}

	if err := ctx.Status(status).JSON(er); err != nil {
		return ctx.Status(500).SendString(fmt.Sprintf("%s", err))
	}
	return nil
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"error"`
}
