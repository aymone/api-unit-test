package main

import (
	"log"
	"net/http"

	"github.com/aymone/api-unit-test/step06/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", api.Router()))
}
