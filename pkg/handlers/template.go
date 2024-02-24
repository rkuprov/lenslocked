package handlers

import "net/http"

type TemplateExecutor interface {
	Execute(w http.ResponseWriter, r *http.Request, data interface{})
}
