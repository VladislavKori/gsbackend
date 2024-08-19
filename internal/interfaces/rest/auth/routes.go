package auth

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest/handler"
)

func Router(userHandler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", userHandler.Register)

	r.Post("/login", userHandler.Login)

	return r
}
