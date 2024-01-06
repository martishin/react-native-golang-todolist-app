package main

import "github.com/gofiber/fiber/v2"

func (application *application) setupRoutes(app *fiber.App) {
	app.Get("/todos", application.getTodos)
	app.Post("/todos", application.createTodo)
}
