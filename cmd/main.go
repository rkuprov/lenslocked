package main

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"

	"lenslocked/cfg"
	"lenslocked/pkg/auth"
	"lenslocked/pkg/handlers"
	"lenslocked/pkg/services"
	"lenslocked/pkg/store"
	"lenslocked/pkg/views"
)

type contact struct {
	Email string
}

func main() {
	ctx := context.Background()
	var c cfg.Cfg
	err := c.Load(filepath.Join("secrets", "cfg.json"))
	if err != nil {
		fmt.Println("error loading config")
		panic(err)
	}
	client, err := store.NewStore(c.Postgres)
	if err != nil {
		panic(err)
	}
	err = client.Setup(ctx)
	if err != nil {
		panic(err)
	}

	csrfMw := csrf.Protect(auth.NewCSRFToken(), csrf.Secure(false))

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

	err = http.ListenAndServe("localhost:3000", csrfMw(r))
	if err != nil {
		panic(err)
	}

}
