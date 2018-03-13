package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("X-Access-Token")
		if accessToken == "password" {
			fmt.Fprint(w, "authenticated with success.\n")
		} else {
			http.Error(w, "you don't have access.", http.StatusForbidden)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
