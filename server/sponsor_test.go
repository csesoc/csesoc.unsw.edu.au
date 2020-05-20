package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestSponsor(t *testing.T) {

	companyName := "Example"
	companyLogo := "https://static.canva.com/static/images/canva_logo_100x100@2x.png"
	companyTier := "100"
	companyExpiry := "2020-11-01T22:08:41+00:00"

	t.Run("Sponsor setup test", func(t *testing.T) {
		resp, err := http.Get("http://localhost:1323/api/sponsors/")
		if err != nil {
			t.Errorf("Could not get perform request.")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusOK)
		var sponsors []*Sponsor
		if err = json.NewDecoder(resp.Body).Decode(&sponsors); err != nil {
			t.Errorf("Error parsing json response: %s", err)
		}
		if len(sponsors) == 0 {
			t.Errorf("Sponsors were not populated.")
		}
	})

	t.Run("Testing sponsor filtering", func(t *testing.T) {
		resp, err := http.Get("http://localhost:1323/api/sponsors/?tier=100")
		if err != nil {
			t.Errorf("Could not perform get sponsors request. Check connection.")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("New sponsor", func(t *testing.T) {
		resp, err := http.PostForm("http://localhost:1323/api/sponsor/", url.Values{
			"name":   {companyName},
			"logo":   {companyLogo},
			"tier":   {companyTier},
			"expiry": {companyExpiry},
		})
		if err != nil {
			t.Errorf("Could not perform post sponsor request. Check connection.")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusCreated)
		var hResp *H
		if err = json.NewDecoder(resp.Body).Decode(&hResp); err != nil {
			t.Errorf("Error parsing json response: %s", err)
		}
	})

	getRequest := "http://localhost:1323/api/sponsor/?name=" + companyName
	t.Run("Get newly created sponsor", func(t *testing.T) {
		resp, err := http.Get(getRequest)
		if err != nil {
			t.Errorf("Could not perform get sponsor request. Check connection.")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusOK)

		var newSponsor *Sponsor
		if err = json.NewDecoder(resp.Body).Decode(&newSponsor); err != nil {
			t.Errorf("Error parsing json response: %s", err)
		} else {
			assertResponseBody(t, newSponsor.Name, companyName)
			assertResponseBody(t, newSponsor.Logo, companyLogo)
			assertResponseBody(t, strconv.Itoa(newSponsor.Tier), companyTier)
		}
	})

	t.Run("Delete newly created sponsor", func(t *testing.T) {
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", getRequest, nil)
		if err != nil {
			t.Errorf("Could not create delete request for sponsor.")
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Could not perform delete sponsor request. Check connection.")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("Check newly removed sponsor", func(t *testing.T) {
		resp, err := http.Get(getRequest)
		if err != nil {
			t.Errorf("Could not perform get sponsor request. Check connection.")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusNotFound)
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong, got %s, want %s", got, want)
	}
}
