package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest/auth"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", auth.Router())

	return r
}
