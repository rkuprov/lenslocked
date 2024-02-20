package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"lenslocked/pkg/views"
)

type contact struct {
	Email string
}
type link struct {
	Title, URL string
}

func main() {
	r := chi.NewRouter()
	r.Get("/", views.StaticView[any](views.Must(views.ParseTemplate("home.gohtml")), link{
		Title: "Contact me!",
		URL:   "localhost:8080/contact",
	}))
	r.Get("/contact", views.StaticView(views.Must(views.ParseTemplate("contact.gohtml")), contact{Email: "kuprov@gmail.com"}))

	http.ListenAndServe(":8080", r)
}
