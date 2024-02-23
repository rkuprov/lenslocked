package main

import (
	"context"
	"fmt"
	"html/template"
	cfg2 "lenslocked/cfg"
	"lenslocked/pkg/services"
	"lenslocked/pkg/store"
	"net/http"
	"path/filepath"

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
	ctx := context.Background()
	var cfg cfg2.Cfg
	err := cfg.Load(filepath.Join("secrets", "cfg.json"))
	if err != nil {
		fmt.Println("error loading config")
		panic(err)
	}
	client, err := store.NewStore(cfg.Postgres)
	if err != nil {
		panic(err)
	}
	err = client.Setup(ctx)
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Get("/", views.StaticView(views.Must(views.ParseTemplate("tailwind.gohtml", "home.gohtml"))))
	r.Get("/contact", views.RenderedView(views.Must(views.ParseTemplate("tailwind.gohtml", "contact.gohtml")), contact{Email: "kuprov@gmail.com"}))

	var u handlers.User
	u.Templates.New = views.Must(views.ParseTemplate("tailwind.gohtml", "signup.gohtml"))
	u.Templates.SignInStatic = views.Must(views.ParseTemplate("tailwind.gohtml", "signin.gohtml"))
	u.Service = services.NewUserService(ctx, client)
	r.Get("/signup", u.New)
	r.Post("/users", u.Create)
	r.Get("/signin", u.SignInStatic)
	r.Post("/signin", u.SignIn)

	err = http.ListenAndServe("localhost:3000", r)
	if err != nil {
		panic(err)
	}

}
