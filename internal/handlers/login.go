package handlers

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseGlob("templates/**/*.html"))

	err := tmpl.ExecuteTemplate(w, "login/index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
