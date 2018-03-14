package api

import "net/http"

type Handler struct {
	AuthService Authenticator
}

type Authenticator interface {
	Auth(string) bool
}

func Router() http.Handler {
	r := http.NewServeMux()

	h := Handler{
		AuthService: &AuthService{},
	}

	r.HandleFunc("/api", h.MainHandler)

	return r
}
