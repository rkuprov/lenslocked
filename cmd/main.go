package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	"lenslocked/pkg/views"
)

type contact struct {
	Email string
}
type link struct {
	Title string
	URL   template.HTML
}

func main() {
	r := chi.NewRouter()
	r.Get("/", views.StaticView[any](views.Must(views.ParseTemplate("tailwind.gohtml", "home.gohtml")), link{
		Title: "Contact me!",
		URL:   "/contact",
	}))
	r.Get("/contact", views.StaticView(views.Must(views.ParseTemplate("tailwind.gohtml", "contact.gohtml")), contact{Email: "kuprov@gmail.com"}))
	r.Get("/signup", views.StaticView[any](views.Must(views.ParseTemplate("tailwind.gohtml", "signup.gohtml")), nil))

	http.ListenAndServe("localhost:3000", r)

}
