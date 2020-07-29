package faq

import (
	"net/http"
	"testing"

	. "csesoc.unsw.edu.au/m/v2/server"
)

func TestFaq(t *testing.T) {
	t.Run("Correct status test", func(t *testing.T) {
		resp, err := http.Get(BASE_URL + FAQ_URL)
		if err != nil {
			t.Errorf("Could not get perform request: ", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusOK)
	})
}
