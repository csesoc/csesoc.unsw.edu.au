package main

import (
	"net/http"
	"testing"
)

func TestSponsor(t *testing.T) {
	t.Run("Sponsor setup test", func(t *testing.T) {
		resp, err := http.Get("http://localhost:1323/api/sponsors/")
		if err != nil {
			t.Errorf("Could not get perform request.")
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})

	t.Run("Testing sponsor filtering", func(t *testing.T) {
		resp, err := http.Get("http://localhost:1323/api/sponsors/?tier=100")
		if err != nil {
			t.Errorf("Could not get perform request.")
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}
