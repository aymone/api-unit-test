package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/api-unit-test/api"
)

func TestMainHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		accessTokenHeader := "password"

		request, err := http.NewRequest("GET", "/api", nil)
		if err != nil {
			t.Fatalf("index request error: %s", err)
		}
		request.Header.Set("X-Access-Token", accessTokenHeader)

		responseWriterMock := httptest.NewRecorder()
		api.MainHandler(responseWriterMock, request)

		expectedCode := http.StatusOK
		if expectedCode != responseWriterMock.Code {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedCode, responseWriterMock.Code)
		}

		expectedBody := []byte("authenticated with success.\n")
		if !bytes.Equal(expectedBody, responseWriterMock.Body.Bytes()) {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedBody, responseWriterMock.Body.String())
		}
	})
}
