package mailing

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	. "csesoc.unsw.edu.au/m/v2/server"
)

const mailingRequestUrl = BASE_URL + MAILING_URL

func TestJoinMessages(t *testing.T) {
	t.Run("Join enquiries", func(t *testing.T) {
		enquiry1 := Enquiry{
			Name:  "Sergio",
			Email: "smr1@gmail.com",
			Body:  "This is the first enquiry",
		}
		enquiry2 := Enquiry{
			Name:  "Sergio",
			Email: "smr2@gmail.com",
			Body:  "This is the second enquiry",
		}
		bundle := []Enquiry{enquiry1, enquiry2}

		composedMsg := joinEnquiries(bundle)
		expectedMsg := "<hr />" +
			"<h3>Enquiry from Sergio &lt;smr1@gmail.com&gt;</h3>" +
			"<p>This is the first enquiry</p>" +
			"<hr />" +
			"<h3>Enquiry from Sergio &lt;smr2@gmail.com&gt;</h3>" +
			"<p>This is the second enquiry</p>" +
			"<hr />"

		if composedMsg != expectedMsg {
			t.Errorf("Output doesn't match expected output\n=== Composed ===\n%v\n=== Expected ===\n%v", composedMsg, expectedMsg)
		}
	})

	t.Run("Join feedbacks", func(t *testing.T) {
		feedback1 := Feedback{
			Name:  "Sergio",
			Email: "smr1@gmail.com",
			Body:  "This is the first feedback",
		}
		feedback2 := Feedback{
			Name:  "Sergio",
			Email: "",
			Body:  "This is the second feedback",
		}
		feedback3 := Feedback{
			Name:  "",
			Email: "smr3@gmail.com",
			Body:  "This is the third feedback",
		}
		bundle := []Feedback{feedback1, feedback2, feedback3}

		composedMsg := joinFeedbacks(bundle)
		expectedMsg := "<hr />" +
			"<p>This is the first feedback</p>" +
			"<i>Feedback from Sergio &lt;smr1@gmail.com&gt;</i>" +
			"<hr />" +
			"<p>This is the second feedback</p>" +
			"<i>Feedback from Sergio</i>" +
			"<hr />" +
			"<p>This is the third feedback</p>" +
			"<i>Feedback from &lt;smr3@gmail.com&gt;</i>" +
			"<hr />"

		if composedMsg != expectedMsg {
			t.Errorf("Output doesn't match expected output\n=== Composed ===\n%v\n=== Expected ===\n%v", composedMsg, expectedMsg)
		}
	})
}

func TestEnquirySuccessful(t *testing.T) {
	formCorrectData := url.Values{
		"name":  {"John Smith"},
		"email": {"john.smith@company.com.au"},
		"body":  {"I'd like to sponsor CSESoc"},
	}

	t.Run("Handle successful sponsorship enquiry", func(t *testing.T) {
		resp, err := http.PostForm(mailingRequestUrl+"/sponsorship", formCorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusAccepted)
	})

	t.Run("Handle successful general enquiry", func(t *testing.T) {
		resp, err := http.PostForm(mailingRequestUrl+"/general", formCorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusAccepted)
	})
}

func TestEnquiryUnsuccessful(t *testing.T) {
	t.Run("Handle request with missing name", func(t *testing.T) {
		formIncorrectData := url.Values{
			"name":  {""},
			"email": {"john.smith@company.com.au"},
			"body":  {"I'd like to sponsor CSESoc"},
		}

		resp, err := http.PostForm(mailingRequestUrl+"/sponsorship", formIncorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})

	t.Run("Handle request with missing email", func(t *testing.T) {
		formIncorrectData := url.Values{
			"name":  {"John Smith"},
			"email": {""},
			"body":  {"I'd like to sponsor CSESoc"},
		}

		resp, err := http.PostForm(mailingRequestUrl+"/sponsorship", formIncorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})

	t.Run("Handle request with no body", func(t *testing.T) {
		formIncorrectData := url.Values{
			"name":  {"John Smith"},
			"email": {"john.smith@company.com.au"},
			"body":  {""},
		}

		resp, err := http.PostForm(mailingRequestUrl+"/sponsorship", formIncorrectData)
		if err != nil {
			t.Errorf("Could not perform POST request: %v", err)
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})

	// http://softwaretesterfriend.com/manual-testing/valid-invalid-email-address-format-validation/
	// Used this website to come up with invalid emails
	invalidEmails := [16]string{
		"example.com",
		"A@b@c@domain.com",
		"a”b(c)d,e:f;gi[j]l@domain.com",
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

			resp, err := http.PostForm(mailingRequestUrl+"/sponsorship", formIncorrectData)
			if err != nil {
				t.Errorf("Could not perform POST request: %v", err)
				return
			}
			defer resp.Body.Close()

			AssertStatus(t, resp.StatusCode, http.StatusBadRequest)
		})
	}
}

func TestFeedbackSuccessful(t *testing.T) {
	t.Run("Feedback Successful", func(t *testing.T) {
		resp, err := http.PostForm(mailingRequestUrl+"/feedback", url.Values{
			"name":  {"John Smith"},
			"email": {"johnsmith@gmail.com"},
			"body":  {"feedback message"},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusAccepted)
	})

	t.Run("Feedback missing name, missing email", func(t *testing.T) {
		resp, err := http.PostForm(mailingRequestUrl+"/feedback", url.Values{
			"body": {"feedback message"},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusAccepted)
	})
}

func TestFeedbackError(t *testing.T) {
	t.Run("Feedback missing name, missing email", func(t *testing.T) {
		resp, err := http.PostForm(mailingRequestUrl+"/feedback", url.Values{
			"email": {"abcde"},
			"body":  {"feedback message"},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})

	t.Run("Feedback missing body", func(t *testing.T) {
		resp, err := http.PostForm(mailingRequestUrl+"/feedback", url.Values{
			"name": {"John Smith"},
		})
		if err != nil {
			t.Errorf("could not perform POST request")
			return
		}
		defer resp.Body.Close()

		AssertStatus(t, resp.StatusCode, http.StatusBadRequest)
	})
}
