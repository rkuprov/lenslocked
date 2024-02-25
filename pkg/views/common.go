package views

import (
	"fmt"
	"html/template"
	"lenslocked/pkg/views/templates"
	"net/http"

	"github.com/gorilla/csrf"

	"lenslocked/pkg/handlers"
)

var _ handlers.TemplateExecutor = (*Template)(nil)

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
	tpl := template.New(patterns[0])
	tpl.Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", fmt.Errorf("csrfField not implemented")
		},
	})
	t, err := tpl.ParseFS(templates.FS, patterns...)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %w", err)
	}

	return &Template{t}, nil
}

func (t *Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tmpl, err := t.Clone()
	if err != nil {
		http.Error(w, "failed cloning template", http.StatusInternalServerError)
	}
	tmpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed executing the template", http.StatusInternalServerError)
	}

}

func StaticView(tmpl *Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, r, nil)
	}
}
func RenderedView(tmpl *Template, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, r, data)
	}
}
