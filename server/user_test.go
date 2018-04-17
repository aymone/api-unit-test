// +build unit

package server_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/api-unit-test/domain"
	"github.com/aymone/api-unit-test/server"

	"github.com/stretchr/testify/assert"
)

type userServiceMock struct {
	GetFn      func(id string) (*domain.User, error)
	GetFnCount int

	InsertFn      func(user *domain.User) (*domain.User, error)
	InsertFnCount int
}

func (us *userServiceMock) Get(id string) (*domain.User, error) {
	us.GetFnCount++
	return us.GetFn(id)
}

func (us *userServiceMock) Insert(user *domain.User) (*domain.User, error) {
	us.InsertFnCount++
	return us.InsertFn(user)
}

func TestUserHandler(t *testing.T) {
	t.Run("get ok", func(t *testing.T) {
		userServiceMock := &userServiceMock{
			GetFn: func(id string) (*domain.User, error) {
				return &domain.User{
					ID: id,
				}, nil
			},
		}
		handler := server.NewHandler(userServiceMock)

		srv := httptest.NewServer(server.New(handler))
		defer srv.Close()

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/users?id=%d", srv.URL, 1), nil)
		assert.NoError(t, err)

		client := &http.Client{}
		res, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, userServiceMock.GetFnCount)
	})

	t.Run("get bad request without id", func(t *testing.T) {
		handler := server.NewHandler(nil)

		srv := httptest.NewServer(server.New(handler))
		defer srv.Close()

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", srv.URL), nil)
		assert.NoError(t, err)

		client := &http.Client{}
		res, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("post created", func(t *testing.T) {
		expectedUserId := "2"
		userServiceMock := &userServiceMock{
			InsertFn: func(user *domain.User) (*domain.User, error) {
				assert.NotNil(t, user)
				assert.Equal(t, expectedUserId, user.ID)

				return user, nil
			},
		}
		handler := server.NewHandler(userServiceMock)

		srv := httptest.NewServer(server.New(handler))
		defer srv.Close()

		jsonBody := []byte(`{"ID": "2"}`)
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
		assert.Equal(t, 1, userServiceMock.InsertFnCount)
	})

	t.Run("post conflict", func(t *testing.T) {
		expectedUserId := "2"
		userServiceMock := &userServiceMock{
			InsertFn: func(user *domain.User) (*domain.User, error) {
				assert.NotNil(t, user)
				assert.Equal(t, expectedUserId, user.ID)

				return nil, domain.ErrDuplicatedKey
			},
		}
		handler := server.NewHandler(userServiceMock)

		srv := httptest.NewServer(server.New(handler))
		defer srv.Close()

		jsonBody := []byte(`{"ID": "2"}`)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/users", srv.URL),
			bytes.NewBuffer(jsonBody),
		)
		assert.NoError(t, err)

		client := &http.Client{}
		res, err := client.Do(req)

		assert.Equal(t, http.StatusConflict, res.StatusCode)
		assert.Equal(t, 1, userServiceMock.InsertFnCount)
	})

}
