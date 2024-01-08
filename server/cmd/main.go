package main

import (
	"flag"
	"fmt"
	"os"

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

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	var dsn string
	flag.StringVar(&dsn, "dsn", fmt.Sprintf("host=%s port=5432 user=postgres password=postgres dbname=todos "+
		"sslmode=disable timezone=UTC connect_timeout=5", dbHost), "Postgres connection string")
	flag.Parse()

	todoRepo := postgresql.NewPostgresTodoRepo(dsn)
	todoService := service.NewTodoService(todoRepo)

	application := &application{todoService: todoService}
	application.setupRoutes(fiberApp)

	if err := fiberApp.Listen(":3000"); err != nil {
		return
	}
}
