package resources

import (
	"encoding/json"
	"net/http"
	"testing"

	. "csesoc.unsw.edu.au/m/v2/server"
)

const resourceRequestURL = BASE_URL + RESOURCES_URL

func TestResources(t *testing.T) {
	t.Run("Preview resources test", func(t *testing.T) {
		resp, err := http.Get(resourceRequestURL + "/preview")
		if err != nil {
			t.Errorf("Could not perform GET request.")
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusOK)
		var previews []*Resource
		if err = json.NewDecoder(resp.Body).Decode(&previews); err != nil {
			t.Errorf("Error parsing json response: %s", err)
		}
	})
}
