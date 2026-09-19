package handlers

import (
	pages "ajaycalicut17/expense-management-go/internal/templ/pages/register"
	"net/http"
)

func IndexRegister(w http.ResponseWriter, r *http.Request) {

	err := pages.IndexRegister().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	// TODO: register
}
