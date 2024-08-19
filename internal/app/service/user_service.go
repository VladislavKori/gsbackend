package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vladislavkori/gsbackend/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      repository.UserRepository
	jwtSecret []byte
}

func NewUserService(repo repository.UserRepository, jwtSecret []byte) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *UserService) RegisterUser(email, password string) (string, int64, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", 0, err
	}

	if user != nil {
		return "", 0, errors.New("user with this email already exists")
	}

	// Устанавливаем время истечения токена
	expirationTime := time.Now().Add(15 * time.Minute)

	// Создаем клеймы
	type Claims struct {
		email      string    `json:"email"`
		created_at time.Time `json:"created_at"`
		jwt.RegisteredClaims
	}

	claims := &Claims{
		email:      email,
		created_at: time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "your_app_name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", 0, err
	}

	// Шифруем пароль
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", 0, err
	}

	// Создаём пользователя
	id, err := s.repo.CreateUser(email, string(hashPassword))
	if err != nil {
		return "", 0, err
	}

	return tokenString, *id, nil
}

func (s *UserService) LoginUser(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Устанавливаем время истечения токена
	expirationTime := time.Now().Add(15 * time.Minute)

	// Создаем клеймы
	type Claims struct {
		email      string    `json:"email"`
		created_at time.Time `json:"created_at"`
		jwt.RegisteredClaims
	}

	claims := &Claims{
		email:      user.Email,
		created_at: time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "your_app_name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
