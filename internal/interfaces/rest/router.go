package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vladislavkori/gsbackend/internal/domain/repository"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest/auth"
)

func Router(userRep repository.UserRepository) http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", auth.Router(userRep))

	return r
}
