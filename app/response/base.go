package response

import "github.com/gofiber/fiber/v2"

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"perPage"`
	TotalPages int   `json:"totalPages"`
	TotalItems int64 `json:"totalItems"`
}

type PaginationResponse struct {
	Data interface{}    `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

func ErrorResponse(c *fiber.Ctx, status int, err error, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error":   err.Error(),
		"message": message,
		"success": false,
	})
}

func SuccessResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"success": true,
	})
}

func SuccessDataResponse(c *fiber.Ctx, status int, data interface{}, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
		"success": true,
	})
}

func SuccessPaginationResponse(c *fiber.Ctx, status int, data interface{}, meta interface{}, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
		"meta":    meta,
		"success": true,
	})
}

func ValidationErrorResponse(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"errors":  errors,
		"message": "Invalid input information",
		"success": false,
	})
}
