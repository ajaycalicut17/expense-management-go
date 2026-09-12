package handlers

import (
	pages "ajaycalicut17/expense-management-go/internal/templ/pages/login"
	"net/http"
)

func IndexLogin(w http.ResponseWriter, r *http.Request) {

	err := pages.IndexLogin().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: login
}
