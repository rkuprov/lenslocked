package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("error getting the current working directory: %v", err)
		return
	}

	t, err := template.ParseFiles(filepath.Join(pwd, "pkg", "templates", "home.gohtml"))
	if err != nil {
		log.Printf("error parsing a template: %v", err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("error executing a template: %v", err)
	}
}
