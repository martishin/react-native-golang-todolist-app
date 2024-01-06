package repository

import (
	"errors"

	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Connection() *gorm.DB
	Todos() []models.Todo
	TodoByID(id int) (*models.Todo, error)
	InsertTodo(m *models.Todo) error
	UpdateTodo(m *models.Todo) error
	DeleteTodo(m *models.Todo) error
}

var ErrTodoAlreadyExists = errors.New("todo already exists")
