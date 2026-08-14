package repositories

import (
	"github.com/mdev4648/todo-api/internal/database"
	"github.com/mdev4648/todo-api/internal/models"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {

	var user models.User

	result := database.DB.
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {

	result := database.DB.Create(user)

	return result.Error
}
