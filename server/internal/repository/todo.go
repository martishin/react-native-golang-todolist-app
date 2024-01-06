package repository

import (
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/repository/postgresql"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Connection() *gorm.DB
	Todos() []models.Todo
	InsertTodo(m *models.Todo) error
}

func NewPostgresTodoRepo(dsn string) *postgresql.PostgresTodoRepo {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		panic("failed to apply migrations")
	}

	return &postgresql.PostgresTodoRepo{DB: db}
}
