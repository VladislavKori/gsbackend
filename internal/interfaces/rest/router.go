package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest/auth"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest/handler"
)

func Router(userHandler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", auth.Router(userHandler))

	return r
}
