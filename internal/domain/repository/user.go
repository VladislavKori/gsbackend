package repository

import "github.com/vladislavkori/gsbackend/internal/domain/entity"

type UserRepository interface {
	// GetUserByID(id int64) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreateUser(email string, password string) (*int64, error)
}
