package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"lenslocked/pkg/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", handlers.Home)
	r.Get("/contact", handlers.Contact)
	http.ListenAndServe("localhost:8080", r)
}
