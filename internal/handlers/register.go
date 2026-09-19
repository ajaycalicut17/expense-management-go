package handlers

import (
	pages "ajaycalicut17/expense-management-go/internal/templ/pages/register"
	"net/http"
)

type RegisterHandler struct{}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{}
}

func (RegisterHandler) Index(w http.ResponseWriter, r *http.Request) {

	err := pages.IndexRegister().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: register
}
