package handlers

import (
	"ajaycalicut17/expense-management-go/internal/response"
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response.Response{
		Status: "ok",
	})
}
