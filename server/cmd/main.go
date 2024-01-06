package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/repository/postgresql"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/service"
)

type application struct {
	todoService service.TodoService
}

func main() {
	fiberApp := fiber.New()
	fiberApp.Use(cors.New())

	var dsn string
	flag.StringVar(&dsn, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=todos "+
		"sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	todoRepo := postgresql.NewPostgresTodoRepo(dsn)
	todoService := service.NewTodoService(todoRepo)

	application := &application{todoService: todoService}
	application.setupRoutes(fiberApp)

	if err := fiberApp.Listen(":3000"); err != nil {
		return
	}
}
