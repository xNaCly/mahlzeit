package models

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Success bool      `json:"success"`
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    fiber.Map `json:"data,omitempty"`
}
