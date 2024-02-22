package handlers

import (
	"fmt"
	"net/http"
)

type User struct {
	Templates struct {
		New    TemplateExecutor
		Create TemplateExecutor
	}
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
	fmt.Fprintf(w, "Email is %s\n", r.FormValue("email"))
	fmt.Fprintf(w, "Password is %s\n", r.FormValue("password"))
}
