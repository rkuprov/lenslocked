package handlers

import (
	"lenslocked/pkg/auth"
	"lenslocked/pkg/datamodel"
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
	user := services.GetUserCtx(r.Context())
	if user == nil {
		log.Default().Println("could not find user in context. redirecting to signin.")
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	session := u.Session.GetSessionID(user.ID)

	data := struct {
		SessionID int

		*datamodel.User
	}{
		SessionID: session,
		User:      user,
	}

	u.Templates.Me.Execute(w, r, data)
}

func (u User) SignOut(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie(auth.CookieTypeSession)
	if err != nil {
		http.Error(w, "could not find session", http.StatusUnauthorized)
	}
	// remove session cookie
	auth.DeleteCookie(w, auth.CookieTypeSession)
	// remove session from db
	err = u.Session.Delete(token.Value)
	if err != nil {
		http.Error(w, "could not delete session", http.StatusInternalServerError)
	}

	// delete user from context
	ctx := services.DeleteUserCtx(r.Context())
	r = r.WithContext(ctx)

	http.Redirect(w, r, "/", http.StatusFound)
}
