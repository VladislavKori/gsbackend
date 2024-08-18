package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=greenshop password=root port=5432 host=localhost sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
