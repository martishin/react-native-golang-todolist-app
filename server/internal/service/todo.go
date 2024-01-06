package service

import (
	"errors"

	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/repository"
)

type TodoService interface {
	Todos() []models.Todo
	TodoByID(id int) (*models.Todo, error)
	CreateTodo(m *models.Todo) error
	UpdateTodo(m *models.Todo) error
}

type DefaultTodoService struct {
	repo repository.TodoRepository
}

var ErrTodoAlreadyExists = errors.New("todo already exists")
var ErrTodoCreationError = errors.New("could not create todo")

func NewTodoService(repo repository.TodoRepository) *DefaultTodoService {
	return &DefaultTodoService{
		repo: repo,
	}
}

func (s *DefaultTodoService) CreateTodo(todo *models.Todo) error {
	if err := s.repo.InsertTodo(todo); err != nil {
		if errors.Is(err, repository.ErrTodoAlreadyExists) {
			return ErrTodoAlreadyExists
		}

		return ErrTodoCreationError
	}
	return nil
}

func (s *DefaultTodoService) Todos() []models.Todo {
	return s.repo.Todos()
}

func (s *DefaultTodoService) TodoByID(id int) (*models.Todo, error) {
	return s.repo.TodoByID(id)
}

func (s *DefaultTodoService) UpdateTodo(todo *models.Todo) error {
	return s.repo.UpdateTodo(todo)
}
