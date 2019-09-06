package controller

import (
	"html/template"
	"net/http"
)

type BaseController struct {
	QueueController
	// CrudController
}

func (l *BaseController) View(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("view/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
