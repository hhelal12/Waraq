package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	user, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// always return array (even if empty)
	if len(users) == 0 {
		json.NewEncoder(w).Encode([]interface{}{})
		return
	}

	json.NewEncoder(w).Encode(users)
}


