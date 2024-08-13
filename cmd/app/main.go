package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vladislavkori/gsbackend/internal/interfaces/rest"
)

func main() {
	r := chi.NewRouter()

	r.Mount("/api", rest.Router())

	http.ListenAndServe(":8080", r)
}
