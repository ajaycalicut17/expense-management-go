package handlers

import (
	"html/template"
	"net/http"
)

func IndexLoginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("internal/templates/login/index.html"))

	data := struct {
		Title string
	}{
		Title: "Login",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: login
}
