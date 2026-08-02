package handlers

import (
	"html/template"
	"net/http"
)

func IndexLogin(w http.ResponseWriter, r *http.Request) {

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

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: login
}
