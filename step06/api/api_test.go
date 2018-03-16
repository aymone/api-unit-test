package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/api-unit-test/step06/api"
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

func TestMainHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		accessTokenHeader := "password"

		request, err := http.NewRequest("GET", "/api", nil)
		if err != nil {
			t.Fatalf("index request error: %s", err)
		}
		request.Header.Set("X-Access-Token", accessTokenHeader)

		responseWriterMock := httptest.NewRecorder()

		handler := api.Handler{
			AuthService: &api.AuthService{},
		}
		handler.MainHandler(responseWriterMock, request)

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

		handler := api.Handler{
			AuthService: &api.AuthService{},
		}
		handler.MainHandler(responseWriterMock, request)

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
