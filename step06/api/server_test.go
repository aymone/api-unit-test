package api_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/api-unit-test/step06/api"
)

func TestRouter(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(api.Router())
		defer srv.Close()

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api", srv.URL), nil)
		if err != nil {
			t.Fatalf("could not create new GET request: %v", err)
		}

		req.Header.Set("X-Access-Token", "password")

		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not send GET request: %v", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code 200, got %v", res.Status)
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read body: %v", err)
		}

		expectedResponseBody := []byte("authenticated with success.")
		if !bytes.Equal(expectedResponseBody, body) {
			t.Errorf("body didn't match:\n\t%q\n\t%q", expectedResponseBody, string(body))
		}
	})

	t.Run("fail", func(t *testing.T) {
		srv := httptest.NewServer(api.Router())
		defer srv.Close()

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api", srv.URL), nil)
		if err != nil {
			t.Fatalf("could not create new GET request: %v", err)
		}

		req.Header.Set("X-Access-Token", "")

		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not send GET request: %v", err)
		}

		if res.StatusCode != http.StatusForbidden {
			t.Errorf("expected status code 403, got %v", res.Status)
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read body: %v", err)
		}

		expectedResponseBody := []byte("you don't have access.\n")
		if !bytes.Equal(expectedResponseBody, body) {
			t.Errorf("body didn't match:\n\t%q\n\t%q", expectedResponseBody, string(body))
		}
	})
}
