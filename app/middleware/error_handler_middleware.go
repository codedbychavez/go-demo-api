package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return ctx.Status(code).SendString(err.Error())
}