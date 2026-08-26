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

func (r *TodoRepository) FindByUserID(userID uint) ([]models.Todo, error) {

	var todos []models.Todo

	result := database.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}
