package repositories

import (
	"github.com/mdev4648/todo-api/internal/database"
	"github.com/mdev4648/todo-api/internal/models"
)

type TodoRepository struct {
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) Create(todo *models.Todo) error {

	return database.DB.Create(todo).Error
}
