package service

import (
	"errors"
	"sort"

	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/repository"
)

type TodoService interface {
	Todos() []models.Todo
	TodoByID(id int) (*models.Todo, error)
	CreateTodo(m *models.Todo) error
	UpdateTodo(m *models.Todo) error
	DeleteTodo(m *models.Todo) error
}

type DefaultTodoService struct {
	repo repository.TodoRepository
}

var ErrTodoAlreadyExists = errors.New("todo already exists")
var ErrTodoCreationError = errors.New("could not create todo")
var ErrTodoUpdateError = errors.New("could not update todo")
var ErrTodoDeleteError = errors.New("could not update todo")
var ErrTodoFetchError = errors.New("could not fetch todo")

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
	todos := s.repo.Todos()
	sort.Slice(
		todos, func(i, j int) bool {
			return todos[i].ID < todos[j].ID
		},
	)
	return todos
}

func (s *DefaultTodoService) TodoByID(id int) (*models.Todo, error) {
	todo, err := s.repo.TodoByID(id)
	if err != nil {
		return nil, ErrTodoFetchError
	}
	return todo, nil
}

func (s *DefaultTodoService) UpdateTodo(todo *models.Todo) error {
	if err := s.repo.UpdateTodo(todo); err != nil {
		return ErrTodoUpdateError
	}
	return nil
}

func (s *DefaultTodoService) DeleteTodo(todo *models.Todo) error {
	if err := s.repo.DeleteTodo(todo); err != nil {
		return ErrTodoDeleteError
	}
	return nil
}
