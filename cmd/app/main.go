package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/vladislavkori/gsbackend/config"
	"github.com/vladislavkori/gsbackend/internal/app/service"
	"github.com/vladislavkori/gsbackend/internal/domain/repository"
	"github.com/vladislavkori/gsbackend/internal/infrastructure/persistence/postgres"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest/handler"
)

func main() {
	r := chi.NewRouter()

	env, err := config.NewEnv()
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	database, err := postgres.NewPostgresDB(repository.PostgresConfig{
		Host:     env.POSTGRESQL_HOST,
		Port:     env.POSTGRESQL_PORT,
		User:     env.POSTGRESQL_USERNAME,
		Password: env.POSTGRESQL_PASSWORD,
		DB_Name:  env.POISTGRESQL_DB_NAME,
		SSLMode:  env.POSTGRESQL_SLLMODE,
	})
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	userRepository := postgres.NewPostgresUserRepository(database)
	userService := service.NewUserService(userRepository, []byte("jwt-secret"))
	userHnaler := handler.NewUserHandler(userService)

	r.Mount("/api", rest.Router(userHnaler))

	http.ListenAndServe(fmt.Sprintf(":%s", env.SERVER_PORT), r)
}
