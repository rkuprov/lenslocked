package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	"lenslocked/pkg/handlers"
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
	r.Get("/", views.StaticView(views.Must(views.ParseTemplate("tailwind.gohtml", "home.gohtml"))))
	r.Get("/contact", views.RenderedView(views.Must(views.ParseTemplate("tailwind.gohtml", "contact.gohtml")), contact{Email: "kuprov@gmail.com"}))

	var u handlers.User
	u.Templates.New = views.Must(views.ParseTemplate("tailwind.gohtml", "signup.gohtml"))
	r.Get("/signup", u.New)
	r.Post("/users", u.Create)

	http.ListenAndServe("localhost:3000", r)

}
