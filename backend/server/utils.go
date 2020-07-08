package utility

import (
	"bufio"
	"os"
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

// ReadSecret returns a docker secret
func ReadSecret(name string) string {
	file, err := os.Open("/run/secrets/" + name)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Return the first line
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}
