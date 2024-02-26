package handlers

import (
	"lenslocked/pkg/auth"
	"log"
	"net/http"

	"lenslocked/pkg/services"
)

type User struct {
	Templates struct {
		New          TemplateExecutor
		Create       TemplateExecutor
		SignInStatic TemplateExecutor
		SignIn       TemplateExecutor
		Me           TemplateExecutor
	}
	Service *services.UserService
	Session *services.SessionService
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusInternalServerError)
	}
	id, err := u.Service.Create(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// creating a session
	session, err := u.Session.Create(id)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	auth.SetCookie(w, auth.CookieTypeSession, session.Token)
	http.Redirect(w, r, "/user/me", http.StatusFound)
}

func (u User) SignInStatic(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignInStatic.Execute(w, r, data)
}

func (u User) SignIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusInternalServerError)
	}
	id, err := u.Service.Authenticate(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	session, err := u.Session.Create(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	auth.SetCookie(w, auth.CookieTypeSession, session.Token)
	http.Redirect(w, r, "/user/me", http.StatusFound)
}

func (u User) Me(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie(auth.CookieTypeSession)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	user, err := u.Session.GetUserForSession(token.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.Templates.Me.Execute(w, r, user)
}

func (u User) SignOut(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie(auth.CookieTypeSession)
	if err != nil {
		http.Error(w, "could not find session", http.StatusUnauthorized)
	}
	csrf, err := r.Cookie("X-CSRF-Token")
	log.Default().Printf("csrf %v", csrf)

	if err != nil {
		http.Error(w, "could not find csrf token", http.StatusUnauthorized)
	}
	// remove session cookie
	auth.SetCookie(w, auth.CookieTypeSession, "")
	// remove session from db
	err = u.Session.Delete(auth.SHAHash(token.Value))
	if err != nil {
		http.Error(w, "could not delete session", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
