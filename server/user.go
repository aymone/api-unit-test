package server

import (
	"encoding/json"
	"net/http"

	"github.com/aymone/api-unit-test/domain"
)

func (h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		h.get(w, r)
		return

	case "POST":
		h.post(w, r)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := keys[0]
	u, err := h.userService.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	u := &domain.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u, err := h.userService.Insert(u)
	if err != nil {
		if err == domain.ErrDuplicatedKey {
			w.WriteHeader(http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
