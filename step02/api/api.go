package api

import (
	"fmt"
	"net/http"
)

// MainHandler ...
func MainHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("X-Access-Token")
	if accessToken == "password" {
		fmt.Fprint(w, "authenticated with success.\n")
		return
	}

	http.Error(w, "you don't have access.", http.StatusForbidden)
}
