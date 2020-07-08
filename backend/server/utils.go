package utility

import (
	"testing"
)

// H - interface for sending JSON
type H map[string]interface{}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong, got %s, want %s", got, want)
	}
}
