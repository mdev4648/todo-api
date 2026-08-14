// package services

// import (
// 	"errors"

// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/mdev4648/todo-api/internal/database"
// 	"github.com/mdev4648/todo-api/internal/dto"
// 	"github.com/mdev4648/todo-api/internal/models"
// )

// func RegisterUser(req dto.RegisterRequest) (*models.User, error) {

// 	var existingUser models.User

// 	result := database.DB.
// 		Where("email = ?", req.Email).
// 		First(&existingUser)

// 	if result.Error == nil {
// 		return nil, errors.New("email already exists")
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword(
// 		[]byte(req.Password),
// 		bcrypt.DefaultCost,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	user := models.User{
// 		Name:     req.Name,
// 		Email:    req.Email,
// 		Password: string(hashedPassword),
// 	}

// 	if err := database.DB.Create(&user).Error; err != nil {
// 		return nil, err
// 	}

//		return &user, nil
//	}
package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/mdev4648/todo-api/internal/dto"
	"github.com/mdev4648/todo-api/internal/models"
	"github.com/mdev4648/todo-api/internal/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(
	userRepository *repositories.UserRepository,
) *UserService {

	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) RegisterUser(
	req dto.RegisterRequest,
) (*models.User, error) {

	_, err := s.UserRepository.FindByEmail(req.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.UserRepository.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
