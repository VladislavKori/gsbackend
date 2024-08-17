package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/vladislavkori/gsbackend/config"
	"github.com/vladislavkori/gsbackend/internal/infrastructure/persistence/postgresql"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest"
)

func main() {
	r := chi.NewRouter()

	if err := postgresql.ConnectToDB(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	env, err := config.NewEnv()
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	r.Mount("/api", rest.Router())

	http.ListenAndServe(fmt.Sprintf(":%s", env.SERVER_PORT), r)
}
