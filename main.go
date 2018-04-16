package main

import (
	"log"
	"net/http"

	"github.com/aymone/api-unit-test/domain/user"
	"github.com/aymone/api-unit-test/server"
	"github.com/aymone/api-unit-test/storage"
)

func main() {
	session, err := storage.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	userStorage, err := storage.NewUserStorage(session)
	if err != nil {
		log.Fatal(err)
	}

	userService := user.NewService(userStorage)
	handler := server.NewHandler(userService)

	srv := server.New(handler)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
