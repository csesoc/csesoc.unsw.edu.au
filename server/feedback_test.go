package main

import (
	"net/http"
	"net/url"
	"testing"
)

const name = "John Smith"
const email = "johnsmith@gmail.com"
const body = "feedback message"
const invalidEmail = "abcde"

func TestFeedbackSuccessful(t *testing.T) {
	t.Run("Feedback Successful", func(t *testing.T) {
		resp, err := http.PostForm("http://localhost:1323/api/enquiry/", url.Values{
			"name":  {name},
			"email": {email},
			"body":  {body},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("Feedback missing name, missing email", func(t *testing.T) {
		resp, err := http.PostForm("http://localhost:1323/api/enquiry/", url.Values{
			"body": {body},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusOK)
	})
}

func TestFeedbackError(t *testing.T) {
	t.Run("Feedback missing name, missing email", func(t *testing.T) {
		resp, err := http.PostForm("http://localhost:1323/api/enquiry/", url.Values{
			"email": {invalidEmail},
			"body":  {body},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})

	t.Run("Feedback missing body", func(t *testing.T) {
		resp, err := http.PostForm("http://localhost:1323/api/enquiry/", url.Values{
			"name": {name},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
		}
		defer resp.Body.Close()

		assertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}
