package postgresql

import (
	"github.com/tty-monkey/react-native-golang-todolist-app/server/internal/models"
	"gorm.io/gorm"
)

type PostgresTodoRepo struct {
	DB *gorm.DB
}

func (p *PostgresTodoRepo) Connection() *gorm.DB {
	return p.DB
}

func (p *PostgresTodoRepo) Todos() []models.Todo {
	var todos []models.Todo
	p.DB.Find(&todos)

	return todos
}

func (p *PostgresTodoRepo) Todo(id int) (*models.Todo, error) {
	var todo models.Todo
	err := p.DB.Find(&todo, id).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (p *PostgresTodoRepo) InsertTodo(todo *models.Todo) error {
	res := p.DB.Create(todo)
	return res.Error
}
