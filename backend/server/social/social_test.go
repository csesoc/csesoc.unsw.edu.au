package social

import (
	"net/http"
	"testing"

	. "csesoc.unsw.edu.au/m/v2/server"
)

func TestSocial(t *testing.T) {
	t.Run("Correct status test", func(t *testing.T) {
		resp, err := http.Get(BASE_URL + SOCIAL_URL)
		if err != nil {
			t.Errorf("Could not perform GET request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusOK)
	})
}
