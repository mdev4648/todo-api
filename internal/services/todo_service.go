package services

import (
	"github.com/mdev4648/todo-api/internal/dto"
	"github.com/mdev4648/todo-api/internal/models"
	"github.com/mdev4648/todo-api/internal/repositories"
)

type TodoService struct {
	TodoRepository *repositories.TodoRepository
}

func NewTodoService(
	todoRepository *repositories.TodoRepository,
) *TodoService {

	return &TodoService{
		TodoRepository: todoRepository,
	}
}

func (s *TodoService) CreateTodo(
	req dto.CreateTodoRequest,
	userID uint,
) (*models.Todo, error) {

	todo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
	}

	if err := s.TodoRepository.Create(&todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (s *TodoService) GetTodos(userID uint) ([]models.Todo, error) {

	return s.TodoRepository.FindByUserID(userID)
}
