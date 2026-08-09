package handlers

import (
	"html/template"
	"net/http"
)

func IndexLogin(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"internal/templates/layouts/base.html",
		"internal/templates/pages/login/index.html",
	))

	err := tmpl.ExecuteTemplate(w, "login/index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: login
}
