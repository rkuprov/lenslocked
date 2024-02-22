package views

import (
	"fmt"
	"html/template"
	"net/http"

	"lenslocked/pkg/templates"
)

type Template struct {
	*template.Template
}

func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseTemplate(patterns ...string) (*Template, error) {
	t, err := template.ParseFS(templates.FS, patterns...)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %w", err)
	}

	return &Template{t}, nil
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	err := t.Template.Execute(w, data)
	if err != nil {
		http.Error(w, "Something went wrong. If the problem persists, please email", http.StatusInternalServerError)
	}

}

func StaticView(tmpl *Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := tmpl.Template.Execute(w, nil)
		if err != nil {
			http.Error(w, "Something went wrong. If the problem persists, please email", http.StatusInternalServerError)
		}
	}
}
func RenderedView(tmpl *Template, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := tmpl.Template.Execute(w, data)
		if err != nil {
			http.Error(w, "Something went wrong. If the problem persists, please email", http.StatusInternalServerError)
		}
	}
}
