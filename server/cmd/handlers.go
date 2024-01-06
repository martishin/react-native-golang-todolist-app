package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
)

func (application *application) getTodos(c *fiber.Ctx) error {
	todos := application.todoService.Todos()

	return c.JSON(todos)
}

func (application *application) createTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return invalidRequestFormat(c)
	}

	if err := application.todoService.CreateTodo(&todo); err != nil {
		return internalServerError(c)
	}

	return c.JSON(&todo)
}

func (application *application) getTodoByID(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return invalidRequestFormat(c)
	}

	todo, err := application.todoService.Todo(id)
	if err != nil {
		return internalServerError(c)
	}

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found",
		})
	}

	return c.JSON(todo)
}

func invalidRequestFormat(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid request format",
	})
}

func internalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "error",
		"message": "Internal server error",
	})
}
