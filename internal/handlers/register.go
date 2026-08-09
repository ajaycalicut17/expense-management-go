package handlers

import (
	"html/template"
	"net/http"
)

func IndexRegister(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"templates/layouts/base.html",
		"templates/pages/register/index.html",
	))

	err := tmpl.ExecuteTemplate(w, "register/index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: register
}
