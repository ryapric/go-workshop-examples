package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// If you intended to support mutiple HTTP methods on the same path, you could
// have refactored this to be called whatever is more meaningful
func TestGet(t *testing.T) {
	t.Run("getRoot returns 200 OK", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(getRoot))
		defer testServer.Close()

		res, err := http.Get(testServer.URL + "/")
		if err != nil {
			t.Fatal(err)
		}

		want := "200 OK"
		got := res.Status

		if want != got {
			t.Errorf("expected response status '%s' but got '%s'", want, got)
		}
	})

	t.Run("getHealtcheck returns 200 OK", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(getHealthcheck))
		defer testServer.Close()

		res, err := http.Get(testServer.URL + "/healthcheck")
		if err != nil {
			t.Fatal(err)
		}

		want := "200 OK"
		got := res.Status

		if want != got {
			t.Errorf("expected response status '%s' but got '%s'", want, got)
		}
	})
}
