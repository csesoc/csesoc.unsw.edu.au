package main

import (
	"net/http"
	"testing"
)

func TestFaq(t *testing.T) {
	t.Run("Correct status test", func(t *testing.T) {
		resp, err := http.Get("http://localhost:1323/api/v1/faq")
		if err != nil {
			t.Errorf("Could not get perform request.")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusOK)
	})
}
