package service

import (
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
	repositories "github.com/tty-monkey/react-native-golang-todolist-app/server/internal/repository"
)

type TodoService interface {
	Todos() []models.Todo
	Todo(id int) (*models.Todo, error)
	CreateTodo(m *models.Todo) error
}

type DefaultTodoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *DefaultTodoService {
	return &DefaultTodoService{
		repo: repo,
	}
}

func (s *DefaultTodoService) CreateTodo(todo *models.Todo) error {
	return s.repo.InsertTodo(todo)
}

func (s *DefaultTodoService) Todos() []models.Todo {
	return s.repo.Todos()
}

func (s *DefaultTodoService) Todo(id int) (*models.Todo, error) {
	return s.repo.Todo(id)
}
