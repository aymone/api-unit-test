package server

import (
	"net/http"
)

type handler interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

func New(h handler) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/users", h.UserHandler)

	return r
}
