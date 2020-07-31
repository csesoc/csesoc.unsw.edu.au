package sponsor

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	. "csesoc.unsw.edu.au/m/v2/server"
)

const companyName = "Example"
const companyLogo = "https://static.canva.com/static/images/canva_logo_100x100@2x.png"
const companyTier = "2"
const companyDetail = "Example"
const companyURL = "https://www.google.com"
const sponsorRequestURL = BASE_URL + SPONSOR_URL

func TestSponsor(t *testing.T) {
	t.Run("Sponsor setup test", func(t *testing.T) {
		resp, err := http.Get(sponsorRequestURL)
		if err != nil {
			t.Errorf("Could not perform GET request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusOK)
		var sponsors []*Sponsor
		if err = json.NewDecoder(resp.Body).Decode(&sponsors); err != nil {
			t.Errorf("Error parsing JSON response: %v", err)
		}
		if len(sponsors) == 0 {
			t.Errorf("Error population Sponsor list")
		}
	})

	t.Run("Testing sponsor filtering", func(t *testing.T) {
		resp, err := http.Get(sponsorRequestURL + "?tier=2")
		if err != nil {
			t.Errorf("Could not perform GET request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("New sponsor", func(t *testing.T) {
		client := &http.Client{}
		form := url.Values{
			"name":   {companyName},
			"logo":   {companyLogo},
			"tier":   {companyTier},
			"detail": {companyDetail},
			"url":    {companyURL},
		}
		req, _ := http.NewRequest("POST", sponsorRequestURL, strings.NewReader(form.Encode()))
		req.Header.Add("Authorization", AUTH_TOKEN)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusCreated)
	})

	t.Run("Get newly created sponsor", func(t *testing.T) {
		resp, err := http.Get(sponsorRequestURL + "/" + companyName)
		if err != nil {
			t.Errorf("Could not perform GET request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusOK)

		var newSponsor *Sponsor
		if err = json.NewDecoder(resp.Body).Decode(&newSponsor); err != nil {
			t.Errorf("Error parsing JSON response: %v", err)
		} else {
			AssertResponseBody(t, newSponsor.Name, companyName)
			AssertResponseBody(t, newSponsor.Logo, companyLogo)
			AssertResponseBody(t, strconv.Itoa(newSponsor.Tier), companyTier)
		}
	})

	t.Run("Delete newly created sponsor", func(t *testing.T) {
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", sponsorRequestURL+"/"+companyName, nil)
		req.Header.Add("Authorization", AUTH_TOKEN)
		if err != nil {
			t.Errorf("Error crafting DELETE request: %v", err)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Could not perform DELETE request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusNoContent)
	})

	t.Run("Check newly removed sponsor", func(t *testing.T) {
		resp, err := http.Get(sponsorRequestURL + "/" + companyName)
		if err != nil {
			t.Errorf("Could not perform GET request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusNotFound)
	})
}

func TestSponsorError(t *testing.T) {
	t.Run("Duplicate Sponsor sponsor", func(t *testing.T) {
		client := &http.Client{}
		form := url.Values{
			"name":   {companyName},
			"logo":   {companyLogo},
			"tier":   {companyTier},
			"detail": {companyDetail},
			"url":    {companyURL},
		}
		req, _ := http.NewRequest("POST", sponsorRequestURL, strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", AUTH_TOKEN)
		req.PostForm = form
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusCreated)

		req, _ = http.NewRequest("POST", sponsorRequestURL, strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", AUTH_TOKEN)
		req.PostForm = form
		resp, err = client.Do(req)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusConflict)

		req, err = http.NewRequest("DELETE", sponsorRequestURL+"/"+companyName, nil)
		req.Header.Add("Authorization", AUTH_TOKEN)
		if err != nil {
			t.Errorf("Error crafting DELETE request: %v", err)
			return
		}
		resp, err = client.Do(req)
		if err != nil {
			t.Errorf("Could not perform DELETE request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusNoContent)
	})

	t.Run("Missing parameters when creating", func(t *testing.T) {
		client := &http.Client{}
		req, _ := http.NewRequest("POST", sponsorRequestURL, nil)
		req.Header.Add("Authorization", AUTH_TOKEN)
		form := url.Values{
			"name": {companyName},
			"logo": {companyLogo},
		}
		req.PostForm = form
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})

	t.Run("Get non existent sponsor", func(t *testing.T) {
		resp, err := http.Get(sponsorRequestURL + "nonexistent")
		if err != nil {
			t.Errorf("Could not perform GET request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusNotFound)
	})
}
