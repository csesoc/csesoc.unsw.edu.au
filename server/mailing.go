package main

import (
	"net/http"

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
var infoEmail string
var sponsorshipBundle []Message
var sponsorshipEmail string

// Mailjet session variables
var publicKey string = "8afb96baef07230483a2a5ceca97d55d"
var secretKey string = "424ad90f25487e6be369a1cbb2a34694"
var mailjetClient *mailjet.Client

// InitMailClient initialises a session with the Mailjet API and stores it in a global variable
func InitMailClient() {
	mailjetClient = mailjet.NewMailjetClient(publicKey, secretKey)
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
		case "sponsorship@csesoc.org.au":
			sponsorshipBundle = append(sponsorshipBundle, message)
		case "info@csesoc.org.au":
			infoBundle = append(infoBundle, message)
		}

		return c.JSON(http.StatusAccepted, H{})
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
		// Dump bundle on txt file
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
