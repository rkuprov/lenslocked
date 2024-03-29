package main

import (
	"context"
	"fmt"
	"github.com/gorilla/csrf"
	"lenslocked/pkg/auth"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"lenslocked/cfg"
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

	sessionSvc := services.NewSessionService(ctx, client)
	userSvc := services.NewUserService(ctx, client)

	u := handlers.User{
		Service: userSvc,
		Session: sessionSvc,
	}

	r := chi.NewRouter()
	csrfMwr := csrf.Protect(auth.NewCSRFToken(), csrf.Secure(false))
	userMwr := services.UserMiddleware{Session: sessionSvc}
	r.Use(csrfMwr)
	r.Use(userMwr.SetUser)

	r.Get("/", views.StaticView(views.Must(views.ParseTemplate("tailwind.gohtml", "home.gohtml"))))
	r.Get("/contact", views.RenderedView(views.Must(views.ParseTemplate("tailwind.gohtml", "contact.gohtml")), contact{Email: "kuprov@gmail.com"}))

	u.Templates.New = views.Must(views.ParseTemplate("tailwind.gohtml", "signup.gohtml"))
	u.Templates.SignInStatic = views.Must(views.ParseTemplate("tailwind.gohtml", "signin.gohtml"))
	u.Templates.Me = views.Must(views.ParseTemplate("tailwind.gohtml", "me.gohtml"))
	r.Get("/signup", u.New)
	r.Post("/users", u.Create)
	r.Get("/signin", u.SignInStatic)
	r.Post("/signin", u.SignIn)
	r.Post("/signout", u.SignOut)

	r.Route("/user", func(r chi.Router) {
		r.Use(csrfMwr)
		r.Use(userMwr.SetUser)
		r.Use(userMwr.RequireUser)
		r.Get("/me", u.Me)
		// todo: try removing to see if csrf cooke breaks
		r.Post("/signout", u.SignOut)
	})

	//middleware := func(in http.Handler) http.Handler {
	//	c := csrf.Protect(auth.NewCSRFToken(), csrf.Secure(false))
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		showRequest(r)
	//c(in).ServeHTTP(w, r)
	//})
	//}

	err = http.ListenAndServe("localhost:3000", r)
	if err != nil {
		panic(err)
	}
}

//func showRequest(r *http.Request) {
//	log.Default().Println(r.Method)
//	for _, cookie := range r.Cookies() {
//		log.Default().Printf("cookie: %v", cookie)
//	}
//	for _, header := range r.Header {
//		log.Default().Printf("header: %s", header)
//	}
//	log.Default().Println("*****************")
//}
