package postgresql

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresTodoRepo struct {
	DB *gorm.DB
}

func NewPostgresTodoRepo(dsn string) *PostgresTodoRepo {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		panic("failed to apply migrations")
	}

	return &PostgresTodoRepo{DB: db}
}

func (p *PostgresTodoRepo) Connection() *gorm.DB {
	return p.DB
}

func (p *PostgresTodoRepo) Todos() []models.Todo {
	var todos []models.Todo
	p.DB.Find(&todos)

	return todos
}

func (p *PostgresTodoRepo) TodoByID(id int) (*models.Todo, error) {
	var todo models.Todo
	err := p.DB.Find(&todo, id).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (p *PostgresTodoRepo) InsertTodo(todo *models.Todo) error {
	err := p.DB.Create(todo).Error

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		//nolint:gocritic // left for illustration purposes
		switch pgErr.Code {
		case "23505":
			return repository.ErrTodoAlreadyExists
		}
	}

	return nil
}

func (p *PostgresTodoRepo) UpdateTodo(todo *models.Todo) error {
	res := p.DB.Save(todo)
	return res.Error
}

func (p *PostgresTodoRepo) DeleteTodo(todo *models.Todo) error {
	res := p.DB.Delete(todo)
	return res.Error
}
