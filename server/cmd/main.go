package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/repository"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/service"
)

type application struct {
	service service.TodoService
}

func main() {
	fiberApp := fiber.New()

	var dsn string
	flag.StringVar(&dsn, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=todos "+
		"sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	todoRepo := repository.NewPostgresTodoRepo(dsn)
	todoService := service.NewTodoService(todoRepo)

	application := &application{service: todoService}
	application.setupRoutes(fiberApp)

	if err := fiberApp.Listen(":3000"); err != nil {
		return
	}
}