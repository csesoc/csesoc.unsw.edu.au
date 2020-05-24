package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mailjet/mailjet-apiv3-go"
)

// Message - struct to contain email message data
type Message struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Body  string `validate:"required"`
}

// Message bundles
var infoBundle []Message
var infoEmail string = "info@csesoc.org.au"
var sponsorshipBundle []Message
var sponsorshipEmail string = "sponsorship@csesoc.org.au"

// Mailjet session variables
var publicKey string = "8afb96baef07230483a2a5ceca97d55d"
var secretKey string = "424ad90f25487e6be369a1cbb2a34694"
var mailjetClient *mailjet.Client

// InitMailClient initialises a session with the Mailjet API and stores it in a global variable
func InitMailClient() {
	mailjetClient = mailjet.NewMailjetClient(publicKey, secretKey)

	// Start mailing timer
	go mailingTimer()
}

// HandleEnquiry by forwarding emails to relevant inboxes
func HandleEnquiry(targetEmail string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract fields from form
		message := Message{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
			Body:  c.FormValue("body"),
		}

		// Validate struct
		if err := c.Validate(message); err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err,
			})
		}

		// Add to bundle
		switch targetEmail {
		case sponsorshipEmail:
			sponsorshipBundle = append(sponsorshipBundle, message)
		case infoEmail:
			infoBundle = append(infoBundle, message)
		}

		return c.JSON(http.StatusAccepted, H{})
	}
}

// This function is executed once in a subroutine and triggers every 15 minutes
func mailingTimer() {
	for {
		time.Sleep(15 * time.Minute)

		go sendEnquiryBundle(infoEmail, infoBundle)
		infoBundle = nil

		go sendEnquiryBundle(sponsorshipEmail, sponsorshipBundle)
		sponsorshipBundle = nil
	}
}

func sendEnquiryBundle(targetEmail string, bundle []Message) {
	// Format message payload
	payload := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "projects.website@csesoc.org.au",
				Name:  "CSESoc Website",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: targetEmail,
				},
			},
			Subject:  "Website enquiry bundle",
			HTMLPart: joinMessages(bundle),
		},
	}

	// Send query
	messages := mailjet.MessagesV31{Info: payload}
	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		// Dump bundle on txt file if it fails to send
	}

}

func joinMessages(bundle []Message) string {
	var message string = ""

	for _, msg := range bundle {
		message += "<hr />"
		message += "<h3>" + "Enquiry from " + msg.Name + " &lt;" + msg.Email + "&gt;" + "</h3>"
		message += "<p>" + msg.Body + "</p>"
	}
	message += "<hr />"

	return message
}
