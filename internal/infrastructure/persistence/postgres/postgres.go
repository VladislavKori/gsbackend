package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vladislavkori/gsbackend/internal/domain/repository"
)

func NewPostgresDB(cfg repository.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf("user=%s dbname=%s password=%s port=%s host=%s sslmode=%s",
			cfg.User, cfg.DB_Name, cfg.Password, cfg.Port, cfg.Host, cfg.SSLMode,
		))

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
