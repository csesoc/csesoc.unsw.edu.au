package main

import (
	"fmt"
	"testing"
)

func TestComposeEmail(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		message := Message{
			Name:  "Sergio",
			Email: "sergio.mercadoruiz9@gmail.com",
			Body:  "This is an enquiry",
		}
		targetEmail := "csesoc@csesoc.org.au"

		composedMsg := composeEmail(message, targetEmail)
		expectedMsg := "Resent-From: " + message.Email + "\r\n" +
			"Reply-To: " + message.Email + "\r\n" +
			"From: " + serverEmail + "\r\n" +
			"To: " + targetEmail + "\r\n" +
			"Subject: " + fmt.Sprintf("Enquiry from '%s' <%s>", message.Name, message.Email) + "\r\n" +
			"\r\n" +
			message.Body

		if composedMsg != expectedMsg {
			t.Errorf("Output doesn't match expected output")
		}
	})
}
