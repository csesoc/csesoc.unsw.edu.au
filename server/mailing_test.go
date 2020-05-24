package main

import (
	"testing"
)

func TestJoinMessages(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		message1 := Message{
			Name:  "Sergio",
			Email: "smr1@gmail.com",
			Body:  "This is the first enquiry",
		}
		message2 := Message{
			Name:  "Sergio",
			Email: "smr2@gmail.com",
			Body:  "This is the second enquiry",
		}
		bundle := []Message{message1, message2}

		composedMsg := joinMessages(bundle)
		expectedMsg := "<hr />" +
			"<h3>Enquiry from Sergio &lt;smr1@gmail.com&gt;</h3>" +
			"<p>This is the first enquiry</p>" +
			"<hr />" +
			"<h3>Enquiry from Sergio &lt;smr2@gmail.com&gt;</h3>" +
			"<p>This is the second enquiry</p>" +
			"<hr />"

		if composedMsg != expectedMsg {
			t.Errorf("Output doesn't match expected output")
		}
	})
}
