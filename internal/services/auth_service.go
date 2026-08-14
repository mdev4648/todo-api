// package services

// import (
// 	"errors"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/mdev4648/todo-api/internal/config"
// 	"github.com/mdev4648/todo-api/internal/database"
// 	"github.com/mdev4648/todo-api/internal/models"
// )

// type AuthService struct {
// 	Config *config.Config
// }

// func NewAuthService(cfg *config.Config) *AuthService {
// 	return &AuthService{
// 		Config: cfg,
// 	}
// }

// func (s *AuthService) Login(email string, password string) (string, *models.User, error) { //receiver Login is belongs to AuthService struct. It takes email and password as input and returns a JWT token, user object and error if any.

// 	var user models.User

// 	result := database.DB.
// 		Where("email = ?", email).
// 		First(&user)

// 	if result.Error != nil {
// 		return "", nil, errors.New("invalid email or password")
// 	}

// 	err := bcrypt.CompareHashAndPassword(
// 		[]byte(user.Password),
// 		[]byte(password),
// 	)

// 	if err != nil {
// 		return "", nil, errors.New("invalid email or password")
// 	}

// 	token := jwt.NewWithClaims(
// 		jwt.SigningMethodHS256,
// 		jwt.MapClaims{
// 			"user_id": user.ID,
// 			"email":   user.Email,
// 			"exp":     time.Now().Add(24 * time.Hour).Unix(),
// 		},
// 	)

// 	signedToken, err := token.SignedString(
// 		[]byte(s.Config.JWTSecret),
// 	)

// 	if err != nil {
// 		return "", nil, err
// 	}

// 	return signedToken, &user, nil
// }

package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/mdev4648/todo-api/internal/config"
	"github.com/mdev4648/todo-api/internal/models"
	"github.com/mdev4648/todo-api/internal/repositories"
)

type AuthService struct {
	Config         *config.Config
	UserRepository *repositories.UserRepository
}

func NewAuthService(
	cfg *config.Config,
	userRepository *repositories.UserRepository,
) *AuthService {

	return &AuthService{
		Config:         cfg,
		UserRepository: userRepository,
	}
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, *models.User, error) {

	user, err := s.UserRepository.FindByEmail(email)

	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp": time.Now().
				Add(24 * time.Hour).
				Unix(),
		},
	)

	signedToken, err := token.SignedString(
		[]byte(s.Config.JWTSecret),
	)

	if err != nil {
		return "", nil, err
	}

	return signedToken, user, nil
}
