package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestSuccessfulEnquiry(t *testing.T) {
	formCorrectData := url.Values{
		"name":  {"John Smith"},
		"email": {"john.smith@company.com.au"},
		"body":  {"I'd like to sponsor CSESoc"},
	}

	t.Run("Handle successful sponsorship enquiry", func(t *testing.T) {
		resp, err := http.PostForm("http://localhost:1323/api/enquiry/sponsorship", formCorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request")
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})

	t.Run("Handle successful general enquiry", func(t *testing.T) {
		resp, err := http.PostForm("http://localhost:1323/api/enquiry/info", formCorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request")
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestUnsuccessfulEnquiry(t *testing.T) {
	t.Run("Handle request with missing name", func(t *testing.T) {
		formIncorrectData := url.Values{
			"name":  {""},
			"email": {"john.smith@company.com.au"},
			"body":  {"I'd like to sponsor CSESoc"},
		}

		resp, err := http.PostForm("http://localhost:1323/api/enquiry/sponsorship", formIncorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request")
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})

	t.Run("Handle request with missing email", func(t *testing.T) {
		formIncorrectData := url.Values{
			"name":  {"John Smith"},
			"email": {""},
			"body":  {"I'd like to sponsor CSESoc"},
		}

		resp, err := http.PostForm("http://localhost:1323/api/enquiry/sponsorship", formIncorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request")
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})

	t.Run("Handle request with no body", func(t *testing.T) {
		formIncorrectData := url.Values{
			"name":  {"John Smith"},
			"email": {"john.smith@company.com.au"},
			"body":  {""},
		}

		resp, err := http.PostForm("http://localhost:1323/api/enquiry/sponsorship", formIncorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request")
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})

	// http://softwaretesterfriend.com/manual-testing/valid-invalid-email-address-format-validation/
	// Used this website to come up with invalid emails
	invalidEmails := [16]string{
		"example.com",
		"A@b@c@domain.com",
		"a”b(c)d,e:f;gi[j]l@domain.com",
		// "abc”test”email@domain.com",
		"abc is”notvalid@domain.com",
		"email.example.com",
		"email@example@example.com",
		".email@example.com",
		"email.@example.com",
		"email..email@example.com",
		"email@example.com (Joe Smith)",
		"email@example",
		"email@-example.com",
		"email@example.web ",
		"email@111.222.333.44444",
		"email@example..com",
		"Abc..123@example.com",
	}
	for index, email := range invalidEmails {
		name := fmt.Sprintf("Handle request with invalid email (%d/%d)", index+1, len(invalidEmails))
		t.Run(name, func(t *testing.T) {
			formIncorrectData := url.Values{
				"name":  {"John Smith"},
				"email": {email},
				"body":  {"I'd like to sponsor CSESoc"},
			}

			resp, err := http.PostForm("http://localhost:1323/api/enquiry/sponsorship", formIncorrectData)
			if err != nil {
				t.Errorf("Could not perform POST request")
			}
			defer resp.Body.Close()

			got := resp.StatusCode
			want := http.StatusBadRequest

			if got != want {
				t.Errorf("got status %d want %d", got, want)
			}
		})
	}
}
