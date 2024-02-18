package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("error getting the current working directory: %v", err)
		return
	}

	t, err := template.ParseFiles(filepath.Join(pwd, "pkg", "templates", "contact.gohtml"))
	if err != nil {
		log.Printf("error parsing a template: %v", err)
		return
	}

	email := struct {
		Email string
	}{
		Email: "kuprov@gmail.com",
	}

	err = t.Execute(w, email)
	if err != nil {
		log.Printf("error executing a template: %v", err)
	}
}
