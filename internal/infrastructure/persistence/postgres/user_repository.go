package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vladislavkori/gsbackend/internal/domain/entity"
	"github.com/vladislavkori/gsbackend/internal/domain/repository"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) repository.UserRepository {
	return &PostgresUserRepository{db: db}
}

// func (r *PostgresUserRepository) GetUserByID(id int64) (*entity.User, error) {
// 	var user *entity.User
// 	query := fmt.Sprintf("SELECT * FROM users WHERE id = %i", id)
// 	rows, err := r.db.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("DB data", rows)
// 	return user, nil
// }

func (r *PostgresUserRepository) FindUserByEmail(email string) (*entity.User, error) {
	var user = &entity.User{}
	query := fmt.Sprintf("SELECT id, email, password, avatar_url, current_delivery_address_id, created_at FROM users WHERE email = '%s'", email)
	err := r.db.DB.QueryRow(query).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.AvatarURL,
		&user.CurrentDeliveryAdressId,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Возвращаем nil и nil ошибку, если пользователь не найден
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) CreateUser(email string, password string) (*int64, error) {
	var userID int64
	query := fmt.Sprintf("insert into users (email, password) values ('%s', '%s') returning id;", email, password)
	err := r.db.DB.QueryRow(query).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}
