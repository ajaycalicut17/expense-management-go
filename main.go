package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseGlob("templates/**/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err := tmpl.ExecuteTemplate(w, "login/index", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
