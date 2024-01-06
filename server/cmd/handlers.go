package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
)

func (application *application) getTodos(c *fiber.Ctx) error {
	todos := application.service.Todos()

	return c.JSON(todos)
}

func (application *application) createTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request format",
		})
	}

	if err := application.service.CreateTodo(&todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
		})
	}

	return c.JSON(&todo)
}
