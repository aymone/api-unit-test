// +build acceptance

package server_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/api-unit-test/domain/user"
	"github.com/aymone/api-unit-test/server"
	"github.com/aymone/api-unit-test/storage"

	"github.com/stretchr/testify/assert"
)

func TestUserAcceptance(t *testing.T) {
	session, err := storage.NewSession()
	if err != nil {
		assert.NoError(t, err)
	}

	userStorage, err := storage.NewUserStorage(session)
	if err != nil {
		assert.NoError(t, err)
	}

	userService := user.NewService(userStorage)
	handler := server.NewHandler(userService)
	srv := httptest.NewServer(server.New(handler))
	defer srv.Close()

	t.Run("get bad request without id", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", srv.URL), nil)
		assert.NoError(t, err)

		client := &http.Client{}
		res, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("post created", func(t *testing.T) {
		jsonBody := []byte(`{"ID": "1"}`)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/users", srv.URL),
			bytes.NewBuffer(jsonBody),
		)
		assert.NoError(t, err)

		client := &http.Client{}
		res, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("get ok", func(t *testing.T) {
		userId := 1
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/users?id=%d", srv.URL, userId), nil)
		assert.NoError(t, err)

		client := &http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("post conflict", func(t *testing.T) {
		jsonBody := []byte(`{"ID": "1"}`)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/users", srv.URL),
			bytes.NewBuffer(jsonBody),
		)
		assert.NoError(t, err)

		client := &http.Client{}
		res, err := client.Do(req)

		assert.Equal(t, http.StatusConflict, res.StatusCode)
	})
}
