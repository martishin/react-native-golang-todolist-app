package main

import "github.com/gofiber/fiber/v2"

func (application *application) setupRoutes(app *fiber.App) {
	app.Get("/todos", application.getTodos)
	app.Post("/todos", application.createTodo)
	app.Get("/todos/:id", application.getTodoByID)
	app.Patch("/todos/:id", application.updateTodoByID)
	app.Delete("/todos/:id", application.deleteTodoByID)
}
