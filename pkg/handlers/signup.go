package handlers

import (
	"fmt"
	"lenslocked/pkg/services"
	"net/http"
)

type User struct {
	Templates struct {
		New    TemplateExecutor
		Create TemplateExecutor
	}
	Service *services.UserService
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
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

	w.Write([]byte(fmt.Sprintf("User created with id: %d", id)))
}
