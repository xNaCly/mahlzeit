package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xnacly/mahlzeit/models"
)

func Ping(c *fiber.Ctx) error {
	return c.JSON(models.ApiResponse{
		Success: true,
		Code:    200,
		Message: "pong",
		Data:    "ping",
	})
}
