package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/api-unit-test/step07/api"
)

func TestAuth(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		auth := &api.AuthService{}
		if !auth.Auth("password") {
			t.Error("authService expected to be true")
		}
	})

	t.Run("fail", func(t *testing.T) {
		auth := &api.AuthService{}
		if auth.Auth("") {
			t.Error("authService expected to be false")
		}
	})
}

type AuthServiceMock struct {
	AuthInvoked bool
	AuthFn      func(token string) bool
}

func (s *AuthServiceMock) Auth(token string) bool {
	s.AuthInvoked = true
	return s.AuthFn(token)
}

func TestMainHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		accessTokenHeader := "password"

		request, err := http.NewRequest("GET", "/api", nil)
		if err != nil {
			t.Fatalf("index request error: %s", err)
		}
		request.Header.Set("X-Access-Token", accessTokenHeader)

		responseWriterMock := httptest.NewRecorder()

		mockedService := AuthServiceMock{
			AuthFn: func(token string) bool {
				return true
			},
		}

		handler := api.Handler{
			AuthService: &mockedService,
		}

		handler.MainHandler(responseWriterMock, request)

		if !mockedService.AuthInvoked {
			t.Error("authServiceMock was expected to be called")
		}

		expectedCode := http.StatusOK
		if expectedCode != responseWriterMock.Code {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedCode, responseWriterMock.Code)
		}

		expectedBody := []byte("authenticated with success.")
		if !bytes.Equal(expectedBody, responseWriterMock.Body.Bytes()) {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedBody, responseWriterMock.Body.String())
		}
	})

	t.Run("fail", func(t *testing.T) {
		accessTokenHeader := ""

		request, err := http.NewRequest("GET", "/api", nil)
		if err != nil {
			t.Fatalf("index request error: %s", err)
		}
		request.Header.Set("X-Access-Token", accessTokenHeader)

		responseWriterMock := httptest.NewRecorder()

		mockedService := AuthServiceMock{
			AuthFn: func(token string) bool {
				return false
			},
		}

		handler := api.Handler{
			AuthService: &mockedService,
		}

		handler.MainHandler(responseWriterMock, request)

		if !mockedService.AuthInvoked {
			t.Error("authServiceMock was expected to be called")
		}

		expectedCode := http.StatusForbidden
		if expectedCode != responseWriterMock.Code {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedCode, responseWriterMock.Code)
		}

		expectedBody := []byte("you don't have access.\n")
		if !bytes.Equal(expectedBody, responseWriterMock.Body.Bytes()) {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedBody, responseWriterMock.Body.String())
		}
	})
}
