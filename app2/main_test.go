package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Server(response, request)

	t.Run("Get server response", func(t *testing.T) {
		got := response.Body.String()
		want := "Hello DevOps Playground!"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

}
