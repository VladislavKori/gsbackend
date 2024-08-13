package auth

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's register route, but now it's not implemented"))
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's login route, but now it's not implemented"))
	})

	return r
}
