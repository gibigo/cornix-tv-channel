package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewHTTPError(ctx *fiber.Ctx, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	if err := ctx.Status(status).JSON(er); err != nil {
		ctx.Status(500).SendString(fmt.Sprintf("%s", err))
		return
	}
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"error"`
}
